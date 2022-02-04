// Package application collects the structures and functions that deal
// with application workloads on k8s
package application

import (
	"context"
	"fmt"
	"regexp"
	"sync"

	"github.com/epinio/epinio/helpers/kubernetes"
	"github.com/epinio/epinio/helpers/kubernetes/tailer"
	"github.com/epinio/epinio/helpers/tracelog"
	"github.com/epinio/epinio/internal/duration"
	"github.com/epinio/epinio/internal/helmchart"
	"github.com/epinio/epinio/internal/namespaces"
	"github.com/epinio/epinio/internal/s3manager"
	"github.com/epinio/epinio/pkg/api/core/v1/models"
	"github.com/pkg/errors"

	epinioappv1 "github.com/epinio/application/api/v1"
	epinioerrors "github.com/epinio/epinio/internal/errors"
	apibatchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/selection"
)

const EpinioApplicationAreaLabel = "epinio.suse.org/area"

// Create generates a new kube app resource in the namespace of the
// namespace. Note that this is the passive resource holding the
// app's configuration. It is not the active workload
func Create(ctx context.Context, cluster *kubernetes.Cluster, app models.AppRef, username string, routes []string) error {
	client, err := cluster.ClientApp()
	if err != nil {
		return err
	}

	// we create the appCRD in the namespace
	obj := &epinioappv1.App{
		Spec: epinioappv1.AppSpec{
			Routes: routes,
			Origin: epinioappv1.AppOrigin{},
		},
	}

	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return err
	}
	us := &unstructured.Unstructured{Object: u}
	us.SetAPIVersion("application.epinio.io/v1")
	us.SetKind("App")
	us.SetName(app.Name)

	_, err = client.Namespace(app.Namespace).Create(ctx, us, metav1.CreateOptions{})
	return err
}

// Get returns the application resource from the cluster.  This should be
// changed to return a typed application struct, like epinioappv1.App if
// needed in the future.
func Get(ctx context.Context, cluster *kubernetes.Cluster, app models.AppRef) (*unstructured.Unstructured, error) {
	client, err := cluster.ClientApp()
	if err != nil {
		return nil, err
	}

	return client.Namespace(app.Namespace).Get(ctx, app.Name, metav1.GetOptions{})
}

// Exists checks if the named application exists or not, and returns an appropriate boolean flag
func Exists(ctx context.Context, cluster *kubernetes.Cluster, app models.AppRef) (bool, error) {
	_, err := Get(ctx, cluster, app)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CurrentlyStaging returns true if there is an active Job for this application.
func CurrentlyStaging(ctx context.Context, cluster *kubernetes.Cluster, namespace, appName string) (bool, error) {

	// Check all jobs for the app for activity.
	selector := fmt.Sprintf("app.kubernetes.io/name=%s,app.kubernetes.io/part-of=%s",
		appName, namespace)

	jobList, err := cluster.ListJobs(ctx, helmchart.StagingNamespace, selector)
	if err != nil {
		return false, err
	}

	staging := false
	for _, job := range jobList.Items {
		for _, condition := range job.Status.Conditions {
			if !(condition.Status == v1.ConditionTrue &&
				(condition.Type == apibatchv1.JobFailed ||
					condition.Type == apibatchv1.JobComplete)) {
				staging = true
			}
		}
	}

	return staging, nil
}

// Lookup locates the named application (and namespace).
func Lookup(ctx context.Context, cluster *kubernetes.Cluster, namespace, appName string) (*models.App, error) {
	meta := models.NewAppRef(appName, namespace)

	ok, err := Exists(ctx, cluster, meta)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	app := meta.App()

	err = fetch(ctx, cluster, app)
	if err != nil {
		app.StatusMessage = err.Error()
		app.Status = models.ApplicationError
	} else {
		err = calculateStatus(ctx, cluster, app)
		if err != nil {
			return app, err
		}
	}

	return app, nil
}

// ListAppRefs returns an app reference for every application resource in the specified
// namespace. If no namespace is specified (empty string) then apps across all namespaces are
// returned.
func ListAppRefs(ctx context.Context, cluster *kubernetes.Cluster, namespace string) ([]models.AppRef, error) {
	client, err := cluster.ClientApp()
	if err != nil {
		return nil, err
	}

	list, err := client.Namespace(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	apps := make([]models.AppRef, 0, len(list.Items))
	for _, app := range list.Items {
		apps = append(apps, models.NewAppRef(app.GetName(), app.GetNamespace()))
	}

	return apps, nil
}

// List returns a list of all available apps in the specified namespace. If no namespace is
// specified (empty string) then apps across all namespaces are returned.
func List(ctx context.Context, cluster *kubernetes.Cluster, namespace string) (models.AppList, error) {

	// Verify namespace, if specified

	if namespace != "" {
		exists, err := namespaces.Exists(ctx, cluster, namespace)
		if err != nil {
			return models.AppList{}, err
		}
		if !exists {
			return models.AppList{}, epinioerrors.NamespaceMissingError{Namespace: namespace}
		}
	}

	// Get references for all apps, deployed or not

	appRefs, err := ListAppRefs(ctx, cluster, namespace)
	if err != nil {
		return models.AppList{}, err
	}

	// Convert references to full application structures

	result := models.AppList{}

	for _, ref := range appRefs {
		app, err := Lookup(ctx, cluster, ref.Namespace, ref.Name)
		if err != nil {
			return result, err
		}
		if app != nil {
			result = append(result, *app)
		}
	}

	return result, nil
}

// Delete removes the named application, its workload (if active), bindings (if any),
// the stored application sources, and any staging jobs from when the application was
// staged (if active). Waits for the application's deployment's pods to disappear
// (if active).
func Delete(ctx context.Context, cluster *kubernetes.Cluster, appRef models.AppRef) error {
	client, err := cluster.ClientApp()
	if err != nil {
		return err
	}

	// delete application resource, will cascade and delete deployment,
	// ingress, service and certificate, environment variables, bindings
	err = client.Namespace(appRef.Namespace).Delete(ctx, appRef.Name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	// delete old staging resources in namespace (helmchart.StagingNamespace)
	err = Unstage(ctx, cluster, appRef, "")
	if err != nil && !apierrors.IsNotFound(err) {
		return err
	}

	// delete staging PVC (the one that holds the "source" and "cache" workspaces)
	err = deleteStagePVC(ctx, cluster, appRef)
	if err != nil && !apierrors.IsNotFound(err) {
		return err
	}

	err = cluster.WaitForPodBySelectorMissing(ctx, nil,
		appRef.Namespace,
		fmt.Sprintf("app.kubernetes.io/name=%s", appRef.Name),
		duration.ToDeployment())
	if err != nil {
		return err
	}

	return nil
}

// deleteStagePVC removes the kube PVC resource which was used to hold the application sources for staging.
func deleteStagePVC(ctx context.Context, cluster *kubernetes.Cluster, appRef models.AppRef) error {
	return cluster.Kubectl.CoreV1().
		PersistentVolumeClaims(helmchart.StagingNamespace).Delete(ctx, appRef.MakePVCName(), metav1.DeleteOptions{})
}

// StageID returns the stage ID of the currently running build, if one exists. It returns an empty string otherwise.
// This method relies on the presence of a workload to get the previous id. There is the case that staging has
// happened, yet there is no workload. Ee.g. by calling the "staging" endpoint but not calling the "deploy"
// endpoint. Since our client doesn't support that scenario, this method doesn't support it either.
func StageID(ctx context.Context, cluster *kubernetes.Cluster, appRef models.AppRef) (string, error) {
	return NewWorkload(cluster, appRef).GetStageID(ctx)
}

// Unstage removes staging resources. It deletes either all Jobs of the
// named application, or all but stageIDCurrent. It also deletes the staged
// objects from the S3 storage except for the current one.
func Unstage(ctx context.Context, cluster *kubernetes.Cluster, appRef models.AppRef, stageIDCurrent string) error {
	s3ConnectionDetails, err := s3manager.GetConnectionDetails(ctx, cluster,
		helmchart.StagingNamespace, helmchart.S3ConnectionDetailsSecretName)
	if err != nil {
		return errors.Wrap(err, "fetching the S3 connection details from the Kubernetes secret")
	}
	s3m, err := s3manager.New(s3ConnectionDetails)
	if err != nil {
		return errors.Wrap(err, "creating an S3 manager")
	}

	jobs, err := cluster.ListJobs(ctx, helmchart.StagingNamespace,
		fmt.Sprintf("app.kubernetes.io/name=%s,app.kubernetes.io/part-of=%s",
			appRef.Name, appRef.Namespace))

	if err != nil {
		return err
	}

	var currentJob *apibatchv1.Job
	for i, job := range jobs.Items {
		id := job.Labels[models.EpinioStageIDLabel]
		// stageIDCurrent is either empty or the id to keep
		if stageIDCurrent != "" && stageIDCurrent == id {
			currentJob = &jobs.Items[i]
			continue
		}

		err := cluster.DeleteJob(ctx, job.ObjectMeta.Namespace, job.ObjectMeta.Name)
		if err != nil {
			return err
		}

		// And the associated secret holding the job environment
		err = cluster.DeleteSecret(ctx, job.ObjectMeta.Namespace, job.ObjectMeta.Name)
		if err != nil {
			return err
		}
	}

	// Cleanup s3 objects
	for _, job := range jobs.Items {
		// skip prs with the same blob as the current one (including the current one)
		if currentJob != nil && job.Labels[models.EpinioStageBlobUIDLabel] == currentJob.Labels[models.EpinioStageBlobUIDLabel] {
			continue
		}

		if err = s3m.DeleteObject(ctx, job.ObjectMeta.Labels[models.EpinioStageBlobUIDLabel]); err != nil {
			return err
		}
	}

	return nil
}

// Logs method writes log lines to the specified logChan. The caller can stop
// the logging with the ctx cancelFunc. It's also the callers responsibility
// to close the logChan when done.
// When stageID is an empty string, no staging logs are returned. If it is set,
// then only logs from that staging process are returned.
func Logs(ctx context.Context, logChan chan tailer.ContainerLogLine, wg *sync.WaitGroup, cluster *kubernetes.Cluster, follow bool, app, stageID, namespace string) error {
	logger := tracelog.NewLogger().WithName("logs-backend").V(2)
	selector := labels.NewSelector()

	var selectors [][]string
	if stageID == "" {
		selectors = [][]string{
			{"app.kubernetes.io/component", "application"},
			{"app.kubernetes.io/part-of", namespace},
			{"app.kubernetes.io/name", app},
		}
	} else {
		selectors = [][]string{
			{"app.kubernetes.io/component", "staging"},
			{models.EpinioStageIDLabel, stageID},
			{"app.kubernetes.io/part-of", namespace},
		}
	}

	for _, req := range selectors {
		req, err := labels.NewRequirement(req[0], selection.Equals, []string{req[1]})
		if err != nil {
			return err
		}
		selector = selector.Add(*req)
	}

	config := &tailer.Config{
		ContainerQuery:        regexp.MustCompile(".*"),
		ExcludeContainerQuery: regexp.MustCompile("linkerd-(proxy|init)"),
		ContainerState:        "running",
		Exclude:               nil,
		Include:               nil,
		Timestamps:            false,
		Since:                 duration.LogHistory(),
		AllNamespaces:         true,
		LabelSelector:         selector,
		TailLines:             nil,
		Namespace:             "",
		PodQuery:              regexp.MustCompile(".*"),
	}

	if follow {
		logger.Info("stream")
		return tailer.StreamLogs(ctx, logChan, wg, config, cluster)
	}

	logger.Info("fetch")
	return tailer.FetchLogs(ctx, logChan, wg, config, cluster)
}

// fetch is a common helper for Lookup and List. It fetches all
// information about an application from the cluster.
func fetch(ctx context.Context, cluster *kubernetes.Cluster, app *models.App) error {
	// Consider delayed loading, i.e. on first access, or for transfer (API response).
	// Consider objects for the information which hide the defered loading.
	// These could also have the necessary modifier methods.
	// See sibling files scale.go, env.go, services.go, ingresses.go.
	// Defered at the moment, the PR is big enough already.

	desiredRoutes, err := DesiredRoutes(ctx, cluster, app.Meta)
	if err != nil {
		return err
	}

	origin, err := Origin(ctx, cluster, app.Meta)
	if err != nil {
		return err
	}

	environment, err := Environment(ctx, cluster, app.Meta)
	if err != nil {
		return err
	}

	instances, err := Scaling(ctx, cluster, app.Meta)
	if err != nil {
		return err
	}

	services, err := BoundServiceNames(ctx, cluster, app.Meta)
	if err != nil {
		return err
	}

	app.Configuration.Instances = &instances
	app.Configuration.Services = services
	app.Configuration.Environment = environment
	app.Configuration.Routes = desiredRoutes
	app.Origin = origin

	// Check if app is active, and if yes, fill the associated parts.
	// May have to straighten the workload structure a bit further.

	app.Workload, err = NewWorkload(cluster, app.Meta).Get(ctx)
	return err
}

// calculateStatus sets the Status field of the App object.
// To decide what the status value should be, it combines various
// pieces of information, i.e. status of possible staging, presence of
// a workload, etc.
//- If Status is ApplicationError, leave it as it (it was set by "Lookup")
//- If there is an active staging job, app is: ApplicationStaging
//- If there is no active staging job and no workload, app is: ApplicationCreated
//- If there is no active staging job and a workload, app is: ApplicationRunning
func calculateStatus(ctx context.Context, cluster *kubernetes.Cluster, app *models.App) error {
	if app.Status == models.ApplicationError {
		return nil
	}
	staging, err := CurrentlyStaging(ctx, cluster, app.Meta.Namespace, app.Meta.Name)
	if err != nil {
		return err
	}
	if staging {
		app.Status = models.ApplicationStaging
		return nil
	}
	if app.Workload == nil {
		app.Status = models.ApplicationCreated
		return nil
	}

	app.Status = models.ApplicationRunning

	return nil
}

package install_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/epinio/epinio/acceptance/helpers/catalog"
	"github.com/epinio/epinio/acceptance/helpers/epinio"
	"github.com/epinio/epinio/acceptance/helpers/proc"
	"github.com/epinio/epinio/acceptance/testenv"
)

var _ = Describe("<Scenario7> RKE, Private CA", func() {
	var (
		flags        []string
		epinioHelper epinio.Epinio
		appName      = catalog.NewAppName()
		loadbalancer string
		metallbURL   string
		localpathURL string
	)

	BeforeEach(func() {
		epinioHelper = epinio.NewEpinioHelper(testenv.EpinioBinaryPath())

		metallbURL = "https://raw.githubusercontent.com/google/metallb/v0.10.3/manifests/metallb.yaml"
		localpathURL = "https://raw.githubusercontent.com/rancher/local-path-provisioner/v0.0.20/deploy/local-path-storage.yaml"

		flags = []string{
			"--skip-default-namespace",
			"--skip-cert-manager",
			"--tls-issuer=private-ca",
		}
	})

	AfterEach(func() {
		out, err := epinioHelper.Uninstall()
		Expect(err).NotTo(HaveOccurred(), out)
	})

	It("installs with loadbalancer IP and pushes an app with env vars", func() {
		By("Installing MetalLB", func() {
			out, err := proc.RunW("kubectl", "create", "namespace", "metallb-system")
			Expect(err).NotTo(HaveOccurred(), out)

			out, err = proc.RunW("kubectl", "apply", "-f", metallbURL)
			Expect(err).NotTo(HaveOccurred(), out)

			out, err = proc.RunW("kubectl", "apply", "-f", testenv.TestAssetPath("config-metallb-rke.yml"))
			Expect(err).NotTo(HaveOccurred(), out)
		})

		By("Configuring local-path storage", func() {
			out, err := proc.RunW("kubectl", "apply", "-f", localpathURL)
			Expect(err).NotTo(HaveOccurred(), out)

			value := `{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}`
			out, err = proc.RunW("kubectl", "patch", "storageclass", "local-path", "-p", value)
			Expect(err).NotTo(HaveOccurred(), out)
		})

		By("Installing CertManager", func() {
			out, err := epinioHelper.Run("install-cert-manager")
			Expect(err).NotTo(HaveOccurred(), out)
			Expect(out).To(ContainSubstring("CertManager deployed"))

			// Create certificate secret and cluster_issuer
			out, err = proc.RunW("kubectl", "apply", "-f", testenv.TestAssetPath("cluster-issuer-private-ca.yml"))
			Expect(err).NotTo(HaveOccurred(), out)
		})

		By("Installing Epinio", func() {
			out, err := epinioHelper.Install(flags...)
			Expect(err).NotTo(HaveOccurred(), out)
			Expect(out).To(ContainSubstring("Epinio installed."))

			out, err = testenv.PatchEpinio()
			Expect(err).ToNot(HaveOccurred(), out)
		})

		By("Checking Loadbalancer IP", func() {
			out, err := proc.RunW("kubectl", "get", "service", "-n", "traefik", "traefik", "-o", "json")
			Expect(err).NotTo(HaveOccurred(), out)

			status := &testenv.LoadBalancerHostname{}
			err = json.Unmarshal([]byte(out), status)
			Expect(err).NotTo(HaveOccurred())
			Expect(status.Status.LoadBalancer.Ingress).To(HaveLen(1))
			loadbalancer = status.Status.LoadBalancer.Ingress[0].IP
			Expect(loadbalancer).ToNot(BeEmpty())
		})

		// Now create the default org which we skipped because
		// it would fail before patching.
		testenv.EnsureDefaultWorkspace(testenv.EpinioBinaryPath())
		out, err := epinioHelper.Run("target", testenv.DefaultWorkspace)
		Expect(err).ToNot(HaveOccurred(), out)

		By("Pushing an app with Env vars", func() {
			out, err := epinioHelper.Run("apps", "create", appName)
			Expect(err).NotTo(HaveOccurred(), out)

			out, err = epinioHelper.Run("apps", "env", "set", appName, "MYVAR", "myvalue")
			Expect(err).ToNot(HaveOccurred(), out)

			out, err = epinioHelper.Run("push",
				"--name", appName,
				"--path", testenv.AssetPath("sample-app"))
			Expect(err).ToNot(HaveOccurred(), out)

			Eventually(func() string {
				out, err := proc.RunW("kubectl", "get", "deployment", "--namespace", testenv.DefaultWorkspace, appName, "-o", "jsonpath={.spec.template.spec.containers[0].env}")
				Expect(err).ToNot(HaveOccurred(), out)
				return out
			}).Should(MatchRegexp("MYVAR"))

			// Verify cluster_issuer is used
			out, err = proc.RunW("kubectl", "get", "certificate",
				"-n", testenv.DefaultWorkspace, appName, "-o", "jsonpath='{.spec.issuerRef.name}'")
			Expect(err).NotTo(HaveOccurred(), out)
			Expect(out).To(Equal("'private-ca'"))
		})
	})
})
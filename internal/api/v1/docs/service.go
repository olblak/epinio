package docs

//go:generate swagger generate spec

import "github.com/epinio/epinio/pkg/api/core/v1/models"

// swagger:route GET /services service ServiceCatalog
// Return all available Epinio Catalog services.
// responses:
//   200: ServiceCatalogResponse

// swagger:parameters ServiceCatalog
type ServiceCatalogParam struct{}

// swagger:response ServiceCatalogResponse
type ServiceCatalogResponse struct {
	// in: body
	Body models.ServiceCatalogResponse
}

// swagger:route GET /services/{Servicename} service ServiceCatalogShow
// Return details of the named Epinio Catalog `Service``.
// responses:
//   200: ServiceCatalogShowResponse

// swagger:parameters ServiceCatalogShow
type ServiceCatalogShowParam struct {
	// in: path
	Servicename string
}

// swagger:response ServiceCatalogShowResponse
type ServiceCatalogShowResponse struct {
	// in: body
	Body models.ServiceCatalogShowResponse
}

// swagger:route POST /namespaces/{Namespace}/services service ServiceCreate
// Create a named instance of an Epinio Catalog Service in the `Namespace`.
// responses:
//   200: ServiceCreateResponse

// swagger:parameters ServiceCreate
type ServiceCreateParam struct {
	// in: path
	Namespace string
	// in: body
	Configuration models.ServiceCreateRequest
}

// swagger:response ServiceCreateResponse
type ServiceCreateResponse struct {
	// in: body
	Body models.Response
}

// swagger:route GET /namespaces/{Namespace}/services service ServiceList
// Return list of service instances in the `Namespace`.
// responses:
//   200: ServiceListResponse

// swagger:parameters ServiceList
type ServiceListParam struct {
	// in: path
	Namespace string
}

// swagger:response ServiceListResponse
type ServiceListResponse struct {
	// in: body
	Body models.ServiceListResponse
}

// swagger:route GET /namespaces/{Namespace}/services/{Service} service ServiceShow
// Return details of the named `Service` instance in the `Namespace`.
// responses:
//   200: ServiceShowResponse

// swagger:response ServiceShowResponse
type ServiceShowResponse struct {
	// in: body
	Body models.ServiceShowResponse
}

// swagger:parameters ServiceShow
type ServiceShowParam struct {
	// in: path
	Namespace string
	// in: path
	Service string
}

// swagger:route DELETE /namespaces/{Namespace}/services/{Service} service ServiceDelete
// Delete the named `Service` instance in the `Namespace`.
// responses:
//   200: ServiceDeleteResponse

// swagger:parameters ServiceDelete
type ServiceDeleteParam struct {
	// in: path
	Namespace string
	// in: path
	Service string
}

// swagger:response ServiceDeleteResponse
type ServiceDeleteResponse struct {
	// in: body
	Body models.Response
}

// swagger:route POST /namespaces/{Namespace}/services/{Servicename}/bind service ServiceBind
// Bind the named `Servicename` instance in the `Namespace` to an App.
// responses:
//   200: ServiceBindResponse

// swagger:parameters ServiceBind
type ServiceBindParam struct {
	// in: path
	Namespace string
	// in: path
	Servicename string
	// in: body
	Configuration models.ServiceBindRequest
}

// swagger:response ServiceBindResponse
type ServiceBindResponse struct {
	// in: body
	Body models.Response
}

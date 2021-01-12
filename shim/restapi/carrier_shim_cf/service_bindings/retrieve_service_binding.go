// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RetrieveServiceBindingHandlerFunc turns a function with the right signature into a retrieve service binding handler
type RetrieveServiceBindingHandlerFunc func(RetrieveServiceBindingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RetrieveServiceBindingHandlerFunc) Handle(params RetrieveServiceBindingParams) middleware.Responder {
	return fn(params)
}

// RetrieveServiceBindingHandler interface for that can handle valid retrieve service binding params
type RetrieveServiceBindingHandler interface {
	Handle(RetrieveServiceBindingParams) middleware.Responder
}

// NewRetrieveServiceBinding creates a new http.Handler for the retrieve service binding operation
func NewRetrieveServiceBinding(ctx *middleware.Context, handler RetrieveServiceBindingHandler) *RetrieveServiceBinding {
	return &RetrieveServiceBinding{Context: ctx, Handler: handler}
}

/*RetrieveServiceBinding swagger:route GET /service_bindings/{guid} serviceBindings retrieveServiceBinding

Retrieve a Particular Service Binding

curl --insecure -i %s/v2/service_bindings/{guid} -X GET -H 'Authorization: %s'

*/
type RetrieveServiceBinding struct {
	Context *middleware.Context
	Handler RetrieveServiceBindingHandler
}

func (o *RetrieveServiceBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRetrieveServiceBindingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
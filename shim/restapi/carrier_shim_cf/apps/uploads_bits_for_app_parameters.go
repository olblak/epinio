// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUploadsBitsForAppParams creates a new UploadsBitsForAppParams object
// no default values defined in spec.
func NewUploadsBitsForAppParams() UploadsBitsForAppParams {

	return UploadsBitsForAppParams{}
}

// UploadsBitsForAppParams contains all the bound params for the uploads bits for app operation
// typically these are obtained from a http.Request
//
// swagger:parameters uploadsBitsForApp
type UploadsBitsForAppParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A binary zip file containing the application bits.
	  Required: true
	  In: formData
	*/
	Application io.ReadCloser
	/*If true, a new asynchronous job is submitted to persist the bits and the job id is included in the response. The client will need to poll the job's status until persistence is completed successfully. If false, the request will block until the bits are persisted synchronously. Defaults to false.
	  In: formData
	*/
	Async *bool
	/*The guid parameter is used as a part of the request URL: '/v2/apps/:guid/bits'
	  Required: true
	  In: path
	*/
	GUID string
	/*Fingerprints of the application bits that have previously been pushed to Cloud Foundry. Each fingerprint must include the file path, sha1 hash, and file size in bytes. Each fingerprint may include the file mode, which must be an octal string with at least read and write permissions for owners. If a mode is not provided, the default mode of 0744 will be used. Fingerprinted bits MUST exist in the Cloud Foundry resource cache or the request (or job, if async) will fail.
	  In: formData
	*/
	Resources *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUploadsBitsForAppParams() beforehand.
func (o *UploadsBitsForAppParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	application, applicationHeader, err := r.FormFile("application")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "application", err))
	} else if err := o.bindApplication(application, applicationHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.Application = &runtime.File{Data: application, Header: applicationHeader}
	}

	fdAsync, fdhkAsync, _ := fds.GetOK("async")
	if err := o.bindAsync(fdAsync, fdhkAsync, route.Formats); err != nil {
		res = append(res, err)
	}

	rGUID, rhkGUID, _ := route.Params.GetOK("guid")
	if err := o.bindGUID(rGUID, rhkGUID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdResources, fdhkResources, _ := fds.GetOK("resources")
	if err := o.bindResources(fdResources, fdhkResources, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindApplication binds file parameter Application.
//
// The only supported validations on files are MinLength and MaxLength
func (o *UploadsBitsForAppParams) bindApplication(file multipart.File, header *multipart.FileHeader) error {
	return nil
}

// bindAsync binds and validates parameter Async from formData.
func (o *UploadsBitsForAppParams) bindAsync(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("async", "formData", "bool", raw)
	}
	o.Async = &value

	return nil
}

// bindGUID binds and validates parameter GUID from path.
func (o *UploadsBitsForAppParams) bindGUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.GUID = raw

	return nil
}

// bindResources binds and validates parameter Resources from formData.
func (o *UploadsBitsForAppParams) bindResources(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Resources = &raw

	return nil
}
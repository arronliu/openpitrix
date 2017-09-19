// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"AppHub/src/api/swagger/models"
)

// GetV1AppsAppIDOKCode is the HTTP code returned for type GetV1AppsAppIDOK
const GetV1AppsAppIDOKCode int = 200

/*GetV1AppsAppIDOK An app

swagger:response getV1AppsAppIdOK
*/
type GetV1AppsAppIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.App `json:"body,omitempty"`
}

// NewGetV1AppsAppIDOK creates GetV1AppsAppIDOK with default headers values
func NewGetV1AppsAppIDOK() *GetV1AppsAppIDOK {
	return &GetV1AppsAppIDOK{}
}

// WithPayload adds the payload to the get v1 apps app Id o k response
func (o *GetV1AppsAppIDOK) WithPayload(payload *models.App) *GetV1AppsAppIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 apps app Id o k response
func (o *GetV1AppsAppIDOK) SetPayload(payload *models.App) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1AppsAppIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetV1AppsAppIDNotFoundCode is the HTTP code returned for type GetV1AppsAppIDNotFound
const GetV1AppsAppIDNotFoundCode int = 404

/*GetV1AppsAppIDNotFound The app does not exist.

swagger:response getV1AppsAppIdNotFound
*/
type GetV1AppsAppIDNotFound struct {
}

// NewGetV1AppsAppIDNotFound creates GetV1AppsAppIDNotFound with default headers values
func NewGetV1AppsAppIDNotFound() *GetV1AppsAppIDNotFound {
	return &GetV1AppsAppIDNotFound{}
}

// WriteResponse to the client
func (o *GetV1AppsAppIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// GetV1AppsAppIDInternalServerErrorCode is the HTTP code returned for type GetV1AppsAppIDInternalServerError
const GetV1AppsAppIDInternalServerErrorCode int = 500

/*GetV1AppsAppIDInternalServerError An unexpected error occured.

swagger:response getV1AppsAppIdInternalServerError
*/
type GetV1AppsAppIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetV1AppsAppIDInternalServerError creates GetV1AppsAppIDInternalServerError with default headers values
func NewGetV1AppsAppIDInternalServerError() *GetV1AppsAppIDInternalServerError {
	return &GetV1AppsAppIDInternalServerError{}
}

// WithPayload adds the payload to the get v1 apps app Id internal server error response
func (o *GetV1AppsAppIDInternalServerError) WithPayload(payload *models.Error) *GetV1AppsAppIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 apps app Id internal server error response
func (o *GetV1AppsAppIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1AppsAppIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
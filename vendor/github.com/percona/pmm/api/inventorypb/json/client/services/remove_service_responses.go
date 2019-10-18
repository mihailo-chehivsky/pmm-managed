// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// RemoveServiceReader is a Reader for the RemoveService structure.
type RemoveServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveServiceOK creates a RemoveServiceOK with default headers values
func NewRemoveServiceOK() *RemoveServiceOK {
	return &RemoveServiceOK{}
}

/*RemoveServiceOK handles this case with default header values.

A successful response.
*/
type RemoveServiceOK struct {
	Payload interface{}
}

func (o *RemoveServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/Remove][%d] removeServiceOk  %+v", 200, o.Payload)
}

func (o *RemoveServiceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RemoveServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveServiceDefault creates a RemoveServiceDefault with default headers values
func NewRemoveServiceDefault(code int) *RemoveServiceDefault {
	return &RemoveServiceDefault{
		_statusCode: code,
	}
}

/*RemoveServiceDefault handles this case with default header values.

An error response.
*/
type RemoveServiceDefault struct {
	_statusCode int

	Payload *RemoveServiceDefaultBody
}

// Code gets the status code for the remove service default response
func (o *RemoveServiceDefault) Code() int {
	return o._statusCode
}

func (o *RemoveServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/Remove][%d] RemoveService default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveServiceDefault) GetPayload() *RemoveServiceDefaultBody {
	return o.Payload
}

func (o *RemoveServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RemoveServiceBody remove service body
swagger:model RemoveServiceBody
*/
type RemoveServiceBody struct {

	// Unique randomly generated instance identifier. Required.
	ServiceID string `json:"service_id,omitempty"`

	// Remove service with all dependencies.
	Force bool `json:"force,omitempty"`
}

// Validate validates this remove service body
func (o *RemoveServiceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveServiceBody) UnmarshalBinary(b []byte) error {
	var res RemoveServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveServiceDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model RemoveServiceDefaultBody
*/
type RemoveServiceDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this remove service default body
func (o *RemoveServiceDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res RemoveServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

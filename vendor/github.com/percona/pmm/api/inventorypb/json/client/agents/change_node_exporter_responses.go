// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeNodeExporterReader is a Reader for the ChangeNodeExporter structure.
type ChangeNodeExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeNodeExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeNodeExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeNodeExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeNodeExporterOK creates a ChangeNodeExporterOK with default headers values
func NewChangeNodeExporterOK() *ChangeNodeExporterOK {
	return &ChangeNodeExporterOK{}
}

/*ChangeNodeExporterOK handles this case with default header values.

A successful response.
*/
type ChangeNodeExporterOK struct {
	Payload *ChangeNodeExporterOKBody
}

func (o *ChangeNodeExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeNodeExporter][%d] changeNodeExporterOk  %+v", 200, o.Payload)
}

func (o *ChangeNodeExporterOK) GetPayload() *ChangeNodeExporterOKBody {
	return o.Payload
}

func (o *ChangeNodeExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeNodeExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeNodeExporterDefault creates a ChangeNodeExporterDefault with default headers values
func NewChangeNodeExporterDefault(code int) *ChangeNodeExporterDefault {
	return &ChangeNodeExporterDefault{
		_statusCode: code,
	}
}

/*ChangeNodeExporterDefault handles this case with default header values.

An error response.
*/
type ChangeNodeExporterDefault struct {
	_statusCode int

	Payload *ChangeNodeExporterDefaultBody
}

// Code gets the status code for the change node exporter default response
func (o *ChangeNodeExporterDefault) Code() int {
	return o._statusCode
}

func (o *ChangeNodeExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeNodeExporter][%d] ChangeNodeExporter default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeNodeExporterDefault) GetPayload() *ChangeNodeExporterDefaultBody {
	return o.Payload
}

func (o *ChangeNodeExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeNodeExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeNodeExporterBody change node exporter body
swagger:model ChangeNodeExporterBody
*/
type ChangeNodeExporterBody struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangeNodeExporterParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change node exporter body
func (o *ChangeNodeExporterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeNodeExporterBody) validateCommon(formats strfmt.Registry) error {

	if swag.IsZero(o.Common) { // not required
		return nil
	}

	if o.Common != nil {
		if err := o.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeNodeExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeNodeExporterBody) UnmarshalBinary(b []byte) error {
	var res ChangeNodeExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeNodeExporterDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model ChangeNodeExporterDefaultBody
*/
type ChangeNodeExporterDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this change node exporter default body
func (o *ChangeNodeExporterDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeNodeExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeNodeExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeNodeExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeNodeExporterOKBody change node exporter OK body
swagger:model ChangeNodeExporterOKBody
*/
type ChangeNodeExporterOKBody struct {

	// node exporter
	NodeExporter *ChangeNodeExporterOKBodyNodeExporter `json:"node_exporter,omitempty"`
}

// Validate validates this change node exporter OK body
func (o *ChangeNodeExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeNodeExporterOKBody) validateNodeExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.NodeExporter) { // not required
		return nil
	}

	if o.NodeExporter != nil {
		if err := o.NodeExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeNodeExporterOk" + "." + "node_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeNodeExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeNodeExporterOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeNodeExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeNodeExporterOKBodyNodeExporter NodeExporter runs on Generic or Container Node and exposes its metrics.
swagger:model ChangeNodeExporterOKBodyNodeExporter
*/
type ChangeNodeExporterOKBodyNodeExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`
}

// Validate validates this change node exporter OK body node exporter
func (o *ChangeNodeExporterOKBodyNodeExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeNodeExporterOkBodyNodeExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeNodeExporterOkBodyNodeExporterTypeStatusPropEnum = append(changeNodeExporterOkBodyNodeExporterTypeStatusPropEnum, v)
	}
}

const (

	// ChangeNodeExporterOKBodyNodeExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	ChangeNodeExporterOKBodyNodeExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// ChangeNodeExporterOKBodyNodeExporterStatusSTARTING captures enum value "STARTING"
	ChangeNodeExporterOKBodyNodeExporterStatusSTARTING string = "STARTING"

	// ChangeNodeExporterOKBodyNodeExporterStatusRUNNING captures enum value "RUNNING"
	ChangeNodeExporterOKBodyNodeExporterStatusRUNNING string = "RUNNING"

	// ChangeNodeExporterOKBodyNodeExporterStatusWAITING captures enum value "WAITING"
	ChangeNodeExporterOKBodyNodeExporterStatusWAITING string = "WAITING"

	// ChangeNodeExporterOKBodyNodeExporterStatusSTOPPING captures enum value "STOPPING"
	ChangeNodeExporterOKBodyNodeExporterStatusSTOPPING string = "STOPPING"

	// ChangeNodeExporterOKBodyNodeExporterStatusDONE captures enum value "DONE"
	ChangeNodeExporterOKBodyNodeExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *ChangeNodeExporterOKBodyNodeExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, changeNodeExporterOkBodyNodeExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *ChangeNodeExporterOKBodyNodeExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeNodeExporterOk"+"."+"node_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeNodeExporterOKBodyNodeExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeNodeExporterOKBodyNodeExporter) UnmarshalBinary(b []byte) error {
	var res ChangeNodeExporterOKBodyNodeExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeNodeExporterParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeNodeExporterParamsBodyCommon
*/
type ChangeNodeExporterParamsBodyCommon struct {

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Disable this Agent. Can't be used with enabled.
	Disable bool `json:"disable,omitempty"`

	// Enable this Agent. Can't be used with disabled.
	Enable bool `json:"enable,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"remove_custom_labels,omitempty"`
}

// Validate validates this change node exporter params body common
func (o *ChangeNodeExporterParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeNodeExporterParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeNodeExporterParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeNodeExporterParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

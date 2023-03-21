// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// ServiceBindingBindingReader is a Reader for the ServiceBindingBinding structure.
type ServiceBindingBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBindingBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBindingBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewServiceBindingBindingCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewServiceBindingBindingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBindingBindingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewServiceBindingBindingConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewServiceBindingBindingUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceBindingBindingOK creates a ServiceBindingBindingOK with default headers values
func NewServiceBindingBindingOK() *ServiceBindingBindingOK {
	return &ServiceBindingBindingOK{}
}

/*
ServiceBindingBindingOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBindingBindingOK struct {
	Payload *models.ServiceBinding
}

// IsSuccess returns true when this service binding binding o k response has a 2xx status code
func (o *ServiceBindingBindingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service binding binding o k response has a 3xx status code
func (o *ServiceBindingBindingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding binding o k response has a 4xx status code
func (o *ServiceBindingBindingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service binding binding o k response has a 5xx status code
func (o *ServiceBindingBindingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding binding o k response a status code equal to that given
func (o *ServiceBindingBindingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service binding binding o k response
func (o *ServiceBindingBindingOK) Code() int {
	return 200
}

func (o *ServiceBindingBindingOK) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingBindingOK  %+v", 200, o.Payload)
}

func (o *ServiceBindingBindingOK) String() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingBindingOK  %+v", 200, o.Payload)
}

func (o *ServiceBindingBindingOK) GetPayload() *models.ServiceBinding {
	return o.Payload
}

func (o *ServiceBindingBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingBindingCreated creates a ServiceBindingBindingCreated with default headers values
func NewServiceBindingBindingCreated() *ServiceBindingBindingCreated {
	return &ServiceBindingBindingCreated{}
}

/*
ServiceBindingBindingCreated describes a response with status code 201, with default header values.

Created
*/
type ServiceBindingBindingCreated struct {
	Payload *models.ServiceBinding
}

// IsSuccess returns true when this service binding binding created response has a 2xx status code
func (o *ServiceBindingBindingCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service binding binding created response has a 3xx status code
func (o *ServiceBindingBindingCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding binding created response has a 4xx status code
func (o *ServiceBindingBindingCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this service binding binding created response has a 5xx status code
func (o *ServiceBindingBindingCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding binding created response a status code equal to that given
func (o *ServiceBindingBindingCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the service binding binding created response
func (o *ServiceBindingBindingCreated) Code() int {
	return 201
}

func (o *ServiceBindingBindingCreated) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingBindingCreated  %+v", 201, o.Payload)
}

func (o *ServiceBindingBindingCreated) String() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingBindingCreated  %+v", 201, o.Payload)
}

func (o *ServiceBindingBindingCreated) GetPayload() *models.ServiceBinding {
	return o.Payload
}

func (o *ServiceBindingBindingCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingBindingAccepted creates a ServiceBindingBindingAccepted with default headers values
func NewServiceBindingBindingAccepted() *ServiceBindingBindingAccepted {
	return &ServiceBindingBindingAccepted{}
}

/*
ServiceBindingBindingAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ServiceBindingBindingAccepted struct {
	Payload *models.AsyncOperation
}

// IsSuccess returns true when this service binding binding accepted response has a 2xx status code
func (o *ServiceBindingBindingAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service binding binding accepted response has a 3xx status code
func (o *ServiceBindingBindingAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding binding accepted response has a 4xx status code
func (o *ServiceBindingBindingAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this service binding binding accepted response has a 5xx status code
func (o *ServiceBindingBindingAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding binding accepted response a status code equal to that given
func (o *ServiceBindingBindingAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the service binding binding accepted response
func (o *ServiceBindingBindingAccepted) Code() int {
	return 202
}

func (o *ServiceBindingBindingAccepted) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingBindingAccepted  %+v", 202, o.Payload)
}

func (o *ServiceBindingBindingAccepted) String() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingBindingAccepted  %+v", 202, o.Payload)
}

func (o *ServiceBindingBindingAccepted) GetPayload() *models.AsyncOperation {
	return o.Payload
}

func (o *ServiceBindingBindingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingBindingBadRequest creates a ServiceBindingBindingBadRequest with default headers values
func NewServiceBindingBindingBadRequest() *ServiceBindingBindingBadRequest {
	return &ServiceBindingBindingBadRequest{}
}

/*
ServiceBindingBindingBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBindingBindingBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding binding bad request response has a 2xx status code
func (o *ServiceBindingBindingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding binding bad request response has a 3xx status code
func (o *ServiceBindingBindingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding binding bad request response has a 4xx status code
func (o *ServiceBindingBindingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding binding bad request response has a 5xx status code
func (o *ServiceBindingBindingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding binding bad request response a status code equal to that given
func (o *ServiceBindingBindingBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service binding binding bad request response
func (o *ServiceBindingBindingBadRequest) Code() int {
	return 400
}

func (o *ServiceBindingBindingBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingBindingBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBindingBindingBadRequest) String() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingBindingBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBindingBindingBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingBindingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingBindingConflict creates a ServiceBindingBindingConflict with default headers values
func NewServiceBindingBindingConflict() *ServiceBindingBindingConflict {
	return &ServiceBindingBindingConflict{}
}

/*
ServiceBindingBindingConflict describes a response with status code 409, with default header values.

Conflict
*/
type ServiceBindingBindingConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding binding conflict response has a 2xx status code
func (o *ServiceBindingBindingConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding binding conflict response has a 3xx status code
func (o *ServiceBindingBindingConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding binding conflict response has a 4xx status code
func (o *ServiceBindingBindingConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding binding conflict response has a 5xx status code
func (o *ServiceBindingBindingConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding binding conflict response a status code equal to that given
func (o *ServiceBindingBindingConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the service binding binding conflict response
func (o *ServiceBindingBindingConflict) Code() int {
	return 409
}

func (o *ServiceBindingBindingConflict) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingBindingConflict  %+v", 409, o.Payload)
}

func (o *ServiceBindingBindingConflict) String() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingBindingConflict  %+v", 409, o.Payload)
}

func (o *ServiceBindingBindingConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingBindingConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingBindingUnprocessableEntity creates a ServiceBindingBindingUnprocessableEntity with default headers values
func NewServiceBindingBindingUnprocessableEntity() *ServiceBindingBindingUnprocessableEntity {
	return &ServiceBindingBindingUnprocessableEntity{}
}

/*
ServiceBindingBindingUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ServiceBindingBindingUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding binding unprocessable entity response has a 2xx status code
func (o *ServiceBindingBindingUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding binding unprocessable entity response has a 3xx status code
func (o *ServiceBindingBindingUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding binding unprocessable entity response has a 4xx status code
func (o *ServiceBindingBindingUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding binding unprocessable entity response has a 5xx status code
func (o *ServiceBindingBindingUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding binding unprocessable entity response a status code equal to that given
func (o *ServiceBindingBindingUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the service binding binding unprocessable entity response
func (o *ServiceBindingBindingUnprocessableEntity) Code() int {
	return 422
}

func (o *ServiceBindingBindingUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingBindingUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceBindingBindingUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingBindingUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceBindingBindingUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingBindingUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

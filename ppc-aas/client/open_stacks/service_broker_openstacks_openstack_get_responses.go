// Code generated by go-swagger; DO NOT EDIT.

package open_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// ServiceBrokerOpenstacksOpenstackGetReader is a Reader for the ServiceBrokerOpenstacksOpenstackGet structure.
type ServiceBrokerOpenstacksOpenstackGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerOpenstacksOpenstackGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerOpenstacksOpenstackGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerOpenstacksOpenstackGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceBrokerOpenstacksOpenstackGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerOpenstacksOpenstackGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /broker/v1/openstacks/{openstack_id}] serviceBroker.openstacks.openstack.get", response, response.Code())
	}
}

// NewServiceBrokerOpenstacksOpenstackGetOK creates a ServiceBrokerOpenstacksOpenstackGetOK with default headers values
func NewServiceBrokerOpenstacksOpenstackGetOK() *ServiceBrokerOpenstacksOpenstackGetOK {
	return &ServiceBrokerOpenstacksOpenstackGetOK{}
}

/*
ServiceBrokerOpenstacksOpenstackGetOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerOpenstacksOpenstackGetOK struct {
	Payload *models.OpenStackInfo
}

// IsSuccess returns true when this service broker openstacks openstack get o k response has a 2xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker openstacks openstack get o k response has a 3xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks openstack get o k response has a 4xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks openstack get o k response has a 5xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks openstack get o k response a status code equal to that given
func (o *ServiceBrokerOpenstacksOpenstackGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker openstacks openstack get o k response
func (o *ServiceBrokerOpenstacksOpenstackGetOK) Code() int {
	return 200
}

func (o *ServiceBrokerOpenstacksOpenstackGetOK) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}][%d] serviceBrokerOpenstacksOpenstackGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerOpenstacksOpenstackGetOK) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}][%d] serviceBrokerOpenstacksOpenstackGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerOpenstacksOpenstackGetOK) GetPayload() *models.OpenStackInfo {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksOpenstackGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenStackInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksOpenstackGetBadRequest creates a ServiceBrokerOpenstacksOpenstackGetBadRequest with default headers values
func NewServiceBrokerOpenstacksOpenstackGetBadRequest() *ServiceBrokerOpenstacksOpenstackGetBadRequest {
	return &ServiceBrokerOpenstacksOpenstackGetBadRequest{}
}

/*
ServiceBrokerOpenstacksOpenstackGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBrokerOpenstacksOpenstackGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks openstack get bad request response has a 2xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks openstack get bad request response has a 3xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks openstack get bad request response has a 4xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks openstack get bad request response has a 5xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks openstack get bad request response a status code equal to that given
func (o *ServiceBrokerOpenstacksOpenstackGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service broker openstacks openstack get bad request response
func (o *ServiceBrokerOpenstacksOpenstackGetBadRequest) Code() int {
	return 400
}

func (o *ServiceBrokerOpenstacksOpenstackGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}][%d] serviceBrokerOpenstacksOpenstackGetBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerOpenstacksOpenstackGetBadRequest) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}][%d] serviceBrokerOpenstacksOpenstackGetBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerOpenstacksOpenstackGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksOpenstackGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksOpenstackGetNotFound creates a ServiceBrokerOpenstacksOpenstackGetNotFound with default headers values
func NewServiceBrokerOpenstacksOpenstackGetNotFound() *ServiceBrokerOpenstacksOpenstackGetNotFound {
	return &ServiceBrokerOpenstacksOpenstackGetNotFound{}
}

/*
ServiceBrokerOpenstacksOpenstackGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServiceBrokerOpenstacksOpenstackGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks openstack get not found response has a 2xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks openstack get not found response has a 3xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks openstack get not found response has a 4xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks openstack get not found response has a 5xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks openstack get not found response a status code equal to that given
func (o *ServiceBrokerOpenstacksOpenstackGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the service broker openstacks openstack get not found response
func (o *ServiceBrokerOpenstacksOpenstackGetNotFound) Code() int {
	return 404
}

func (o *ServiceBrokerOpenstacksOpenstackGetNotFound) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}][%d] serviceBrokerOpenstacksOpenstackGetNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerOpenstacksOpenstackGetNotFound) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}][%d] serviceBrokerOpenstacksOpenstackGetNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerOpenstacksOpenstackGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksOpenstackGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksOpenstackGetInternalServerError creates a ServiceBrokerOpenstacksOpenstackGetInternalServerError with default headers values
func NewServiceBrokerOpenstacksOpenstackGetInternalServerError() *ServiceBrokerOpenstacksOpenstackGetInternalServerError {
	return &ServiceBrokerOpenstacksOpenstackGetInternalServerError{}
}

/*
ServiceBrokerOpenstacksOpenstackGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerOpenstacksOpenstackGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks openstack get internal server error response has a 2xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks openstack get internal server error response has a 3xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks openstack get internal server error response has a 4xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks openstack get internal server error response has a 5xx status code
func (o *ServiceBrokerOpenstacksOpenstackGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker openstacks openstack get internal server error response a status code equal to that given
func (o *ServiceBrokerOpenstacksOpenstackGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the service broker openstacks openstack get internal server error response
func (o *ServiceBrokerOpenstacksOpenstackGetInternalServerError) Code() int {
	return 500
}

func (o *ServiceBrokerOpenstacksOpenstackGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}][%d] serviceBrokerOpenstacksOpenstackGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerOpenstacksOpenstackGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}][%d] serviceBrokerOpenstacksOpenstackGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerOpenstacksOpenstackGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksOpenstackGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

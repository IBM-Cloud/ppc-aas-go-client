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

// ServiceBrokerOpenstacksHostsGetReader is a Reader for the ServiceBrokerOpenstacksHostsGet structure.
type ServiceBrokerOpenstacksHostsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerOpenstacksHostsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerOpenstacksHostsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerOpenstacksHostsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceBrokerOpenstacksHostsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerOpenstacksHostsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceBrokerOpenstacksHostsGetOK creates a ServiceBrokerOpenstacksHostsGetOK with default headers values
func NewServiceBrokerOpenstacksHostsGetOK() *ServiceBrokerOpenstacksHostsGetOK {
	return &ServiceBrokerOpenstacksHostsGetOK{}
}

/*
ServiceBrokerOpenstacksHostsGetOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerOpenstacksHostsGetOK struct {
	Payload *models.HostInfo
}

// IsSuccess returns true when this service broker openstacks hosts get o k response has a 2xx status code
func (o *ServiceBrokerOpenstacksHostsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker openstacks hosts get o k response has a 3xx status code
func (o *ServiceBrokerOpenstacksHostsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks hosts get o k response has a 4xx status code
func (o *ServiceBrokerOpenstacksHostsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks hosts get o k response has a 5xx status code
func (o *ServiceBrokerOpenstacksHostsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks hosts get o k response a status code equal to that given
func (o *ServiceBrokerOpenstacksHostsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker openstacks hosts get o k response
func (o *ServiceBrokerOpenstacksHostsGetOK) Code() int {
	return 200
}

func (o *ServiceBrokerOpenstacksHostsGetOK) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/hosts/{hostname}][%d] serviceBrokerOpenstacksHostsGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerOpenstacksHostsGetOK) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/hosts/{hostname}][%d] serviceBrokerOpenstacksHostsGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerOpenstacksHostsGetOK) GetPayload() *models.HostInfo {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksHostsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksHostsGetBadRequest creates a ServiceBrokerOpenstacksHostsGetBadRequest with default headers values
func NewServiceBrokerOpenstacksHostsGetBadRequest() *ServiceBrokerOpenstacksHostsGetBadRequest {
	return &ServiceBrokerOpenstacksHostsGetBadRequest{}
}

/*
ServiceBrokerOpenstacksHostsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBrokerOpenstacksHostsGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks hosts get bad request response has a 2xx status code
func (o *ServiceBrokerOpenstacksHostsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks hosts get bad request response has a 3xx status code
func (o *ServiceBrokerOpenstacksHostsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks hosts get bad request response has a 4xx status code
func (o *ServiceBrokerOpenstacksHostsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks hosts get bad request response has a 5xx status code
func (o *ServiceBrokerOpenstacksHostsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks hosts get bad request response a status code equal to that given
func (o *ServiceBrokerOpenstacksHostsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service broker openstacks hosts get bad request response
func (o *ServiceBrokerOpenstacksHostsGetBadRequest) Code() int {
	return 400
}

func (o *ServiceBrokerOpenstacksHostsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/hosts/{hostname}][%d] serviceBrokerOpenstacksHostsGetBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerOpenstacksHostsGetBadRequest) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/hosts/{hostname}][%d] serviceBrokerOpenstacksHostsGetBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerOpenstacksHostsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksHostsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksHostsGetNotFound creates a ServiceBrokerOpenstacksHostsGetNotFound with default headers values
func NewServiceBrokerOpenstacksHostsGetNotFound() *ServiceBrokerOpenstacksHostsGetNotFound {
	return &ServiceBrokerOpenstacksHostsGetNotFound{}
}

/*
ServiceBrokerOpenstacksHostsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServiceBrokerOpenstacksHostsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks hosts get not found response has a 2xx status code
func (o *ServiceBrokerOpenstacksHostsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks hosts get not found response has a 3xx status code
func (o *ServiceBrokerOpenstacksHostsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks hosts get not found response has a 4xx status code
func (o *ServiceBrokerOpenstacksHostsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks hosts get not found response has a 5xx status code
func (o *ServiceBrokerOpenstacksHostsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks hosts get not found response a status code equal to that given
func (o *ServiceBrokerOpenstacksHostsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the service broker openstacks hosts get not found response
func (o *ServiceBrokerOpenstacksHostsGetNotFound) Code() int {
	return 404
}

func (o *ServiceBrokerOpenstacksHostsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/hosts/{hostname}][%d] serviceBrokerOpenstacksHostsGetNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerOpenstacksHostsGetNotFound) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/hosts/{hostname}][%d] serviceBrokerOpenstacksHostsGetNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerOpenstacksHostsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksHostsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksHostsGetInternalServerError creates a ServiceBrokerOpenstacksHostsGetInternalServerError with default headers values
func NewServiceBrokerOpenstacksHostsGetInternalServerError() *ServiceBrokerOpenstacksHostsGetInternalServerError {
	return &ServiceBrokerOpenstacksHostsGetInternalServerError{}
}

/*
ServiceBrokerOpenstacksHostsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerOpenstacksHostsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks hosts get internal server error response has a 2xx status code
func (o *ServiceBrokerOpenstacksHostsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks hosts get internal server error response has a 3xx status code
func (o *ServiceBrokerOpenstacksHostsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks hosts get internal server error response has a 4xx status code
func (o *ServiceBrokerOpenstacksHostsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks hosts get internal server error response has a 5xx status code
func (o *ServiceBrokerOpenstacksHostsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker openstacks hosts get internal server error response a status code equal to that given
func (o *ServiceBrokerOpenstacksHostsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the service broker openstacks hosts get internal server error response
func (o *ServiceBrokerOpenstacksHostsGetInternalServerError) Code() int {
	return 500
}

func (o *ServiceBrokerOpenstacksHostsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/hosts/{hostname}][%d] serviceBrokerOpenstacksHostsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerOpenstacksHostsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks/{openstack_id}/hosts/{hostname}][%d] serviceBrokerOpenstacksHostsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerOpenstacksHostsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksHostsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

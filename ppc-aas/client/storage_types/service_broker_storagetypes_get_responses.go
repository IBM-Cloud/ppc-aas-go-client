// Code generated by go-swagger; DO NOT EDIT.

package storage_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// ServiceBrokerStoragetypesGetReader is a Reader for the ServiceBrokerStoragetypesGet structure.
type ServiceBrokerStoragetypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerStoragetypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerStoragetypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewServiceBrokerStoragetypesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceBrokerStoragetypesGetOK creates a ServiceBrokerStoragetypesGetOK with default headers values
func NewServiceBrokerStoragetypesGetOK() *ServiceBrokerStoragetypesGetOK {
	return &ServiceBrokerStoragetypesGetOK{}
}

/*
ServiceBrokerStoragetypesGetOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerStoragetypesGetOK struct {
	Payload models.StorageTypes
}

// IsSuccess returns true when this service broker storagetypes get o k response has a 2xx status code
func (o *ServiceBrokerStoragetypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker storagetypes get o k response has a 3xx status code
func (o *ServiceBrokerStoragetypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker storagetypes get o k response has a 4xx status code
func (o *ServiceBrokerStoragetypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker storagetypes get o k response has a 5xx status code
func (o *ServiceBrokerStoragetypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker storagetypes get o k response a status code equal to that given
func (o *ServiceBrokerStoragetypesGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServiceBrokerStoragetypesGetOK) Error() string {
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerStoragetypesGetOK) String() string {
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerStoragetypesGetOK) GetPayload() models.StorageTypes {
	return o.Payload
}

func (o *ServiceBrokerStoragetypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerStoragetypesGetInternalServerError creates a ServiceBrokerStoragetypesGetInternalServerError with default headers values
func NewServiceBrokerStoragetypesGetInternalServerError() *ServiceBrokerStoragetypesGetInternalServerError {
	return &ServiceBrokerStoragetypesGetInternalServerError{}
}

/*
ServiceBrokerStoragetypesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerStoragetypesGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker storagetypes get internal server error response has a 2xx status code
func (o *ServiceBrokerStoragetypesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker storagetypes get internal server error response has a 3xx status code
func (o *ServiceBrokerStoragetypesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker storagetypes get internal server error response has a 4xx status code
func (o *ServiceBrokerStoragetypesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker storagetypes get internal server error response has a 5xx status code
func (o *ServiceBrokerStoragetypesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker storagetypes get internal server error response a status code equal to that given
func (o *ServiceBrokerStoragetypesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ServiceBrokerStoragetypesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerStoragetypesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerStoragetypesGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerStoragetypesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

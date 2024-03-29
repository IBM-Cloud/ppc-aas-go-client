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

// ServiceBrokerOpenstacksPostReader is a Reader for the ServiceBrokerOpenstacksPost structure.
type ServiceBrokerOpenstacksPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerOpenstacksPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerOpenstacksPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewServiceBrokerOpenstacksPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerOpenstacksPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewServiceBrokerOpenstacksPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewServiceBrokerOpenstacksPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerOpenstacksPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /broker/v1/openstacks] serviceBroker.openstacks.post", response, response.Code())
	}
}

// NewServiceBrokerOpenstacksPostOK creates a ServiceBrokerOpenstacksPostOK with default headers values
func NewServiceBrokerOpenstacksPostOK() *ServiceBrokerOpenstacksPostOK {
	return &ServiceBrokerOpenstacksPostOK{}
}

/*
ServiceBrokerOpenstacksPostOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerOpenstacksPostOK struct {
	Payload *models.OpenStack
}

// IsSuccess returns true when this service broker openstacks post o k response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker openstacks post o k response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post o k response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks post o k response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks post o k response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker openstacks post o k response
func (o *ServiceBrokerOpenstacksPostOK) Code() int {
	return 200
}

func (o *ServiceBrokerOpenstacksPostOK) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostOK) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostOK) GetPayload() *models.OpenStack {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostCreated creates a ServiceBrokerOpenstacksPostCreated with default headers values
func NewServiceBrokerOpenstacksPostCreated() *ServiceBrokerOpenstacksPostCreated {
	return &ServiceBrokerOpenstacksPostCreated{}
}

/*
ServiceBrokerOpenstacksPostCreated describes a response with status code 201, with default header values.

Created
*/
type ServiceBrokerOpenstacksPostCreated struct {
	Payload *models.OpenStack
}

// IsSuccess returns true when this service broker openstacks post created response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker openstacks post created response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post created response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks post created response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks post created response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the service broker openstacks post created response
func (o *ServiceBrokerOpenstacksPostCreated) Code() int {
	return 201
}

func (o *ServiceBrokerOpenstacksPostCreated) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostCreated  %+v", 201, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostCreated) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostCreated  %+v", 201, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostCreated) GetPayload() *models.OpenStack {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostBadRequest creates a ServiceBrokerOpenstacksPostBadRequest with default headers values
func NewServiceBrokerOpenstacksPostBadRequest() *ServiceBrokerOpenstacksPostBadRequest {
	return &ServiceBrokerOpenstacksPostBadRequest{}
}

/*
ServiceBrokerOpenstacksPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBrokerOpenstacksPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks post bad request response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks post bad request response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post bad request response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks post bad request response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks post bad request response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service broker openstacks post bad request response
func (o *ServiceBrokerOpenstacksPostBadRequest) Code() int {
	return 400
}

func (o *ServiceBrokerOpenstacksPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostBadRequest) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostConflict creates a ServiceBrokerOpenstacksPostConflict with default headers values
func NewServiceBrokerOpenstacksPostConflict() *ServiceBrokerOpenstacksPostConflict {
	return &ServiceBrokerOpenstacksPostConflict{}
}

/*
ServiceBrokerOpenstacksPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type ServiceBrokerOpenstacksPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks post conflict response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks post conflict response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post conflict response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks post conflict response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks post conflict response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the service broker openstacks post conflict response
func (o *ServiceBrokerOpenstacksPostConflict) Code() int {
	return 409
}

func (o *ServiceBrokerOpenstacksPostConflict) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostConflict  %+v", 409, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostConflict) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostConflict  %+v", 409, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostUnprocessableEntity creates a ServiceBrokerOpenstacksPostUnprocessableEntity with default headers values
func NewServiceBrokerOpenstacksPostUnprocessableEntity() *ServiceBrokerOpenstacksPostUnprocessableEntity {
	return &ServiceBrokerOpenstacksPostUnprocessableEntity{}
}

/*
ServiceBrokerOpenstacksPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ServiceBrokerOpenstacksPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks post unprocessable entity response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks post unprocessable entity response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post unprocessable entity response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks post unprocessable entity response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks post unprocessable entity response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the service broker openstacks post unprocessable entity response
func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) Code() int {
	return 422
}

func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostInternalServerError creates a ServiceBrokerOpenstacksPostInternalServerError with default headers values
func NewServiceBrokerOpenstacksPostInternalServerError() *ServiceBrokerOpenstacksPostInternalServerError {
	return &ServiceBrokerOpenstacksPostInternalServerError{}
}

/*
ServiceBrokerOpenstacksPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerOpenstacksPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks post internal server error response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks post internal server error response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post internal server error response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks post internal server error response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker openstacks post internal server error response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the service broker openstacks post internal server error response
func (o *ServiceBrokerOpenstacksPostInternalServerError) Code() int {
	return 500
}

func (o *ServiceBrokerOpenstacksPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

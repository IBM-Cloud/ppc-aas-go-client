// Code generated by go-swagger; DO NOT EDIT.

package hardware_platforms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// ServiceBrokerHardwareplatformsGetReader is a Reader for the ServiceBrokerHardwareplatformsGet structure.
type ServiceBrokerHardwareplatformsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerHardwareplatformsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerHardwareplatformsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewServiceBrokerHardwareplatformsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceBrokerHardwareplatformsGetOK creates a ServiceBrokerHardwareplatformsGetOK with default headers values
func NewServiceBrokerHardwareplatformsGetOK() *ServiceBrokerHardwareplatformsGetOK {
	return &ServiceBrokerHardwareplatformsGetOK{}
}

/*
ServiceBrokerHardwareplatformsGetOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerHardwareplatformsGetOK struct {
	Payload models.HardwarePlatforms
}

// IsSuccess returns true when this service broker hardwareplatforms get o k response has a 2xx status code
func (o *ServiceBrokerHardwareplatformsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker hardwareplatforms get o k response has a 3xx status code
func (o *ServiceBrokerHardwareplatformsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker hardwareplatforms get o k response has a 4xx status code
func (o *ServiceBrokerHardwareplatformsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker hardwareplatforms get o k response has a 5xx status code
func (o *ServiceBrokerHardwareplatformsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker hardwareplatforms get o k response a status code equal to that given
func (o *ServiceBrokerHardwareplatformsGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServiceBrokerHardwareplatformsGetOK) Error() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetOK) String() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetOK) GetPayload() models.HardwarePlatforms {
	return o.Payload
}

func (o *ServiceBrokerHardwareplatformsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerHardwareplatformsGetInternalServerError creates a ServiceBrokerHardwareplatformsGetInternalServerError with default headers values
func NewServiceBrokerHardwareplatformsGetInternalServerError() *ServiceBrokerHardwareplatformsGetInternalServerError {
	return &ServiceBrokerHardwareplatformsGetInternalServerError{}
}

/*
ServiceBrokerHardwareplatformsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerHardwareplatformsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker hardwareplatforms get internal server error response has a 2xx status code
func (o *ServiceBrokerHardwareplatformsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker hardwareplatforms get internal server error response has a 3xx status code
func (o *ServiceBrokerHardwareplatformsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker hardwareplatforms get internal server error response has a 4xx status code
func (o *ServiceBrokerHardwareplatformsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker hardwareplatforms get internal server error response has a 5xx status code
func (o *ServiceBrokerHardwareplatformsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker hardwareplatforms get internal server error response a status code equal to that given
func (o *ServiceBrokerHardwareplatformsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ServiceBrokerHardwareplatformsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /broker/v1/hardware-platforms][%d] serviceBrokerHardwareplatformsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerHardwareplatformsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerHardwareplatformsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

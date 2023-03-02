// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudNetworksPortsGetReader is a Reader for the PcloudNetworksPortsGet structure.
type PcloudNetworksPortsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudNetworksPortsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudNetworksPortsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPcloudNetworksPortsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudNetworksPortsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudNetworksPortsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudNetworksPortsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudNetworksPortsGetOK creates a PcloudNetworksPortsGetOK with default headers values
func NewPcloudNetworksPortsGetOK() *PcloudNetworksPortsGetOK {
	return &PcloudNetworksPortsGetOK{}
}

/*
PcloudNetworksPortsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudNetworksPortsGetOK struct {
	Payload *models.NetworkPort
}

// IsSuccess returns true when this pcloud networks ports get o k response has a 2xx status code
func (o *PcloudNetworksPortsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud networks ports get o k response has a 3xx status code
func (o *PcloudNetworksPortsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports get o k response has a 4xx status code
func (o *PcloudNetworksPortsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks ports get o k response has a 5xx status code
func (o *PcloudNetworksPortsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports get o k response a status code equal to that given
func (o *PcloudNetworksPortsGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudNetworksPortsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksPortsGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksPortsGetOK) GetPayload() *models.NetworkPort {
	return o.Payload
}

func (o *PcloudNetworksPortsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsGetUnauthorized creates a PcloudNetworksPortsGetUnauthorized with default headers values
func NewPcloudNetworksPortsGetUnauthorized() *PcloudNetworksPortsGetUnauthorized {
	return &PcloudNetworksPortsGetUnauthorized{}
}

/*
PcloudNetworksPortsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudNetworksPortsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports get unauthorized response has a 2xx status code
func (o *PcloudNetworksPortsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports get unauthorized response has a 3xx status code
func (o *PcloudNetworksPortsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports get unauthorized response has a 4xx status code
func (o *PcloudNetworksPortsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports get unauthorized response has a 5xx status code
func (o *PcloudNetworksPortsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports get unauthorized response a status code equal to that given
func (o *PcloudNetworksPortsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudNetworksPortsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudNetworksPortsGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudNetworksPortsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsGetForbidden creates a PcloudNetworksPortsGetForbidden with default headers values
func NewPcloudNetworksPortsGetForbidden() *PcloudNetworksPortsGetForbidden {
	return &PcloudNetworksPortsGetForbidden{}
}

/*
PcloudNetworksPortsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudNetworksPortsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports get forbidden response has a 2xx status code
func (o *PcloudNetworksPortsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports get forbidden response has a 3xx status code
func (o *PcloudNetworksPortsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports get forbidden response has a 4xx status code
func (o *PcloudNetworksPortsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports get forbidden response has a 5xx status code
func (o *PcloudNetworksPortsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports get forbidden response a status code equal to that given
func (o *PcloudNetworksPortsGetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudNetworksPortsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudNetworksPortsGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudNetworksPortsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsGetNotFound creates a PcloudNetworksPortsGetNotFound with default headers values
func NewPcloudNetworksPortsGetNotFound() *PcloudNetworksPortsGetNotFound {
	return &PcloudNetworksPortsGetNotFound{}
}

/*
PcloudNetworksPortsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudNetworksPortsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports get not found response has a 2xx status code
func (o *PcloudNetworksPortsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports get not found response has a 3xx status code
func (o *PcloudNetworksPortsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports get not found response has a 4xx status code
func (o *PcloudNetworksPortsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports get not found response has a 5xx status code
func (o *PcloudNetworksPortsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports get not found response a status code equal to that given
func (o *PcloudNetworksPortsGetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudNetworksPortsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudNetworksPortsGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudNetworksPortsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsGetInternalServerError creates a PcloudNetworksPortsGetInternalServerError with default headers values
func NewPcloudNetworksPortsGetInternalServerError() *PcloudNetworksPortsGetInternalServerError {
	return &PcloudNetworksPortsGetInternalServerError{}
}

/*
PcloudNetworksPortsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudNetworksPortsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports get internal server error response has a 2xx status code
func (o *PcloudNetworksPortsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports get internal server error response has a 3xx status code
func (o *PcloudNetworksPortsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports get internal server error response has a 4xx status code
func (o *PcloudNetworksPortsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks ports get internal server error response has a 5xx status code
func (o *PcloudNetworksPortsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud networks ports get internal server error response a status code equal to that given
func (o *PcloudNetworksPortsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudNetworksPortsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksPortsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksPortsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

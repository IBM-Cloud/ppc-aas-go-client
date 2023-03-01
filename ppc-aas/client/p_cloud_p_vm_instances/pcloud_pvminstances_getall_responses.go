// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudPvminstancesGetallReader is a Reader for the PcloudPvminstancesGetall structure.
type PcloudPvminstancesGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesGetallOK creates a PcloudPvminstancesGetallOK with default headers values
func NewPcloudPvminstancesGetallOK() *PcloudPvminstancesGetallOK {
	return &PcloudPvminstancesGetallOK{}
}

/*
PcloudPvminstancesGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesGetallOK struct {
	Payload *models.PVMInstances
}

// IsSuccess returns true when this pcloud pvminstances getall o k response has a 2xx status code
func (o *PcloudPvminstancesGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances getall o k response has a 3xx status code
func (o *PcloudPvminstancesGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances getall o k response has a 4xx status code
func (o *PcloudPvminstancesGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances getall o k response has a 5xx status code
func (o *PcloudPvminstancesGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances getall o k response a status code equal to that given
func (o *PcloudPvminstancesGetallOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudPvminstancesGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesGetallOK) GetPayload() *models.PVMInstances {
	return o.Payload
}

func (o *PcloudPvminstancesGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PVMInstances)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesGetallBadRequest creates a PcloudPvminstancesGetallBadRequest with default headers values
func NewPcloudPvminstancesGetallBadRequest() *PcloudPvminstancesGetallBadRequest {
	return &PcloudPvminstancesGetallBadRequest{}
}

/*
PcloudPvminstancesGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances getall bad request response has a 2xx status code
func (o *PcloudPvminstancesGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances getall bad request response has a 3xx status code
func (o *PcloudPvminstancesGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances getall bad request response has a 4xx status code
func (o *PcloudPvminstancesGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances getall bad request response has a 5xx status code
func (o *PcloudPvminstancesGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances getall bad request response a status code equal to that given
func (o *PcloudPvminstancesGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudPvminstancesGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesGetallUnauthorized creates a PcloudPvminstancesGetallUnauthorized with default headers values
func NewPcloudPvminstancesGetallUnauthorized() *PcloudPvminstancesGetallUnauthorized {
	return &PcloudPvminstancesGetallUnauthorized{}
}

/*
PcloudPvminstancesGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances getall unauthorized response has a 2xx status code
func (o *PcloudPvminstancesGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances getall unauthorized response has a 3xx status code
func (o *PcloudPvminstancesGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances getall unauthorized response has a 4xx status code
func (o *PcloudPvminstancesGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances getall unauthorized response has a 5xx status code
func (o *PcloudPvminstancesGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances getall unauthorized response a status code equal to that given
func (o *PcloudPvminstancesGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudPvminstancesGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesGetallForbidden creates a PcloudPvminstancesGetallForbidden with default headers values
func NewPcloudPvminstancesGetallForbidden() *PcloudPvminstancesGetallForbidden {
	return &PcloudPvminstancesGetallForbidden{}
}

/*
PcloudPvminstancesGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances getall forbidden response has a 2xx status code
func (o *PcloudPvminstancesGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances getall forbidden response has a 3xx status code
func (o *PcloudPvminstancesGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances getall forbidden response has a 4xx status code
func (o *PcloudPvminstancesGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances getall forbidden response has a 5xx status code
func (o *PcloudPvminstancesGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances getall forbidden response a status code equal to that given
func (o *PcloudPvminstancesGetallForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudPvminstancesGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesGetallInternalServerError creates a PcloudPvminstancesGetallInternalServerError with default headers values
func NewPcloudPvminstancesGetallInternalServerError() *PcloudPvminstancesGetallInternalServerError {
	return &PcloudPvminstancesGetallInternalServerError{}
}

/*
PcloudPvminstancesGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances getall internal server error response has a 2xx status code
func (o *PcloudPvminstancesGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances getall internal server error response has a 3xx status code
func (o *PcloudPvminstancesGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances getall internal server error response has a 4xx status code
func (o *PcloudPvminstancesGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances getall internal server error response has a 5xx status code
func (o *PcloudPvminstancesGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances getall internal server error response a status code equal to that given
func (o *PcloudPvminstancesGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudPvminstancesGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances][%d] pcloudPvminstancesGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

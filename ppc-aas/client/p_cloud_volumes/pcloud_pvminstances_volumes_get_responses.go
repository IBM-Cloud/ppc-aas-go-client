// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudPvminstancesVolumesGetReader is a Reader for the PcloudPvminstancesVolumesGet structure.
type PcloudPvminstancesVolumesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesVolumesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesVolumesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesVolumesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesVolumesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesVolumesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesVolumesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesVolumesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesVolumesGetOK creates a PcloudPvminstancesVolumesGetOK with default headers values
func NewPcloudPvminstancesVolumesGetOK() *PcloudPvminstancesVolumesGetOK {
	return &PcloudPvminstancesVolumesGetOK{}
}

/*
PcloudPvminstancesVolumesGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesVolumesGetOK struct {
	Payload *models.Volume
}

// IsSuccess returns true when this pcloud pvminstances volumes get o k response has a 2xx status code
func (o *PcloudPvminstancesVolumesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances volumes get o k response has a 3xx status code
func (o *PcloudPvminstancesVolumesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes get o k response has a 4xx status code
func (o *PcloudPvminstancesVolumesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances volumes get o k response has a 5xx status code
func (o *PcloudPvminstancesVolumesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes get o k response a status code equal to that given
func (o *PcloudPvminstancesVolumesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud pvminstances volumes get o k response
func (o *PcloudPvminstancesVolumesGetOK) Code() int {
	return 200
}

func (o *PcloudPvminstancesVolumesGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetOK) GetPayload() *models.Volume {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetBadRequest creates a PcloudPvminstancesVolumesGetBadRequest with default headers values
func NewPcloudPvminstancesVolumesGetBadRequest() *PcloudPvminstancesVolumesGetBadRequest {
	return &PcloudPvminstancesVolumesGetBadRequest{}
}

/*
PcloudPvminstancesVolumesGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesVolumesGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes get bad request response has a 2xx status code
func (o *PcloudPvminstancesVolumesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes get bad request response has a 3xx status code
func (o *PcloudPvminstancesVolumesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes get bad request response has a 4xx status code
func (o *PcloudPvminstancesVolumesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes get bad request response has a 5xx status code
func (o *PcloudPvminstancesVolumesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes get bad request response a status code equal to that given
func (o *PcloudPvminstancesVolumesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances volumes get bad request response
func (o *PcloudPvminstancesVolumesGetBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesVolumesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetUnauthorized creates a PcloudPvminstancesVolumesGetUnauthorized with default headers values
func NewPcloudPvminstancesVolumesGetUnauthorized() *PcloudPvminstancesVolumesGetUnauthorized {
	return &PcloudPvminstancesVolumesGetUnauthorized{}
}

/*
PcloudPvminstancesVolumesGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesVolumesGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes get unauthorized response has a 2xx status code
func (o *PcloudPvminstancesVolumesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes get unauthorized response has a 3xx status code
func (o *PcloudPvminstancesVolumesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes get unauthorized response has a 4xx status code
func (o *PcloudPvminstancesVolumesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes get unauthorized response has a 5xx status code
func (o *PcloudPvminstancesVolumesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes get unauthorized response a status code equal to that given
func (o *PcloudPvminstancesVolumesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances volumes get unauthorized response
func (o *PcloudPvminstancesVolumesGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesVolumesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetForbidden creates a PcloudPvminstancesVolumesGetForbidden with default headers values
func NewPcloudPvminstancesVolumesGetForbidden() *PcloudPvminstancesVolumesGetForbidden {
	return &PcloudPvminstancesVolumesGetForbidden{}
}

/*
PcloudPvminstancesVolumesGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesVolumesGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes get forbidden response has a 2xx status code
func (o *PcloudPvminstancesVolumesGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes get forbidden response has a 3xx status code
func (o *PcloudPvminstancesVolumesGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes get forbidden response has a 4xx status code
func (o *PcloudPvminstancesVolumesGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes get forbidden response has a 5xx status code
func (o *PcloudPvminstancesVolumesGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes get forbidden response a status code equal to that given
func (o *PcloudPvminstancesVolumesGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud pvminstances volumes get forbidden response
func (o *PcloudPvminstancesVolumesGetForbidden) Code() int {
	return 403
}

func (o *PcloudPvminstancesVolumesGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetNotFound creates a PcloudPvminstancesVolumesGetNotFound with default headers values
func NewPcloudPvminstancesVolumesGetNotFound() *PcloudPvminstancesVolumesGetNotFound {
	return &PcloudPvminstancesVolumesGetNotFound{}
}

/*
PcloudPvminstancesVolumesGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesVolumesGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes get not found response has a 2xx status code
func (o *PcloudPvminstancesVolumesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes get not found response has a 3xx status code
func (o *PcloudPvminstancesVolumesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes get not found response has a 4xx status code
func (o *PcloudPvminstancesVolumesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes get not found response has a 5xx status code
func (o *PcloudPvminstancesVolumesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes get not found response a status code equal to that given
func (o *PcloudPvminstancesVolumesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud pvminstances volumes get not found response
func (o *PcloudPvminstancesVolumesGetNotFound) Code() int {
	return 404
}

func (o *PcloudPvminstancesVolumesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetInternalServerError creates a PcloudPvminstancesVolumesGetInternalServerError with default headers values
func NewPcloudPvminstancesVolumesGetInternalServerError() *PcloudPvminstancesVolumesGetInternalServerError {
	return &PcloudPvminstancesVolumesGetInternalServerError{}
}

/*
PcloudPvminstancesVolumesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesVolumesGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes get internal server error response has a 2xx status code
func (o *PcloudPvminstancesVolumesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes get internal server error response has a 3xx status code
func (o *PcloudPvminstancesVolumesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes get internal server error response has a 4xx status code
func (o *PcloudPvminstancesVolumesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances volumes get internal server error response has a 5xx status code
func (o *PcloudPvminstancesVolumesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances volumes get internal server error response a status code equal to that given
func (o *PcloudPvminstancesVolumesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances volumes get internal server error response
func (o *PcloudPvminstancesVolumesGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesVolumesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesVolumesGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

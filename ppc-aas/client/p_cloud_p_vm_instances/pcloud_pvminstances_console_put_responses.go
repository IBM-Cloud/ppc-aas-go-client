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

// PcloudPvminstancesConsolePutReader is a Reader for the PcloudPvminstancesConsolePut structure.
type PcloudPvminstancesConsolePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesConsolePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesConsolePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesConsolePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesConsolePutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesConsolePutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesConsolePutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesConsolePutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesConsolePutOK creates a PcloudPvminstancesConsolePutOK with default headers values
func NewPcloudPvminstancesConsolePutOK() *PcloudPvminstancesConsolePutOK {
	return &PcloudPvminstancesConsolePutOK{}
}

/*
PcloudPvminstancesConsolePutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesConsolePutOK struct {
	Payload *models.ConsoleLanguage
}

// IsSuccess returns true when this pcloud pvminstances console put o k response has a 2xx status code
func (o *PcloudPvminstancesConsolePutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances console put o k response has a 3xx status code
func (o *PcloudPvminstancesConsolePutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console put o k response has a 4xx status code
func (o *PcloudPvminstancesConsolePutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances console put o k response has a 5xx status code
func (o *PcloudPvminstancesConsolePutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console put o k response a status code equal to that given
func (o *PcloudPvminstancesConsolePutOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudPvminstancesConsolePutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesConsolePutOK) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesConsolePutOK) GetPayload() *models.ConsoleLanguage {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleLanguage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePutBadRequest creates a PcloudPvminstancesConsolePutBadRequest with default headers values
func NewPcloudPvminstancesConsolePutBadRequest() *PcloudPvminstancesConsolePutBadRequest {
	return &PcloudPvminstancesConsolePutBadRequest{}
}

/*
PcloudPvminstancesConsolePutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesConsolePutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console put bad request response has a 2xx status code
func (o *PcloudPvminstancesConsolePutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console put bad request response has a 3xx status code
func (o *PcloudPvminstancesConsolePutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console put bad request response has a 4xx status code
func (o *PcloudPvminstancesConsolePutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console put bad request response has a 5xx status code
func (o *PcloudPvminstancesConsolePutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console put bad request response a status code equal to that given
func (o *PcloudPvminstancesConsolePutBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudPvminstancesConsolePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesConsolePutBadRequest) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesConsolePutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePutUnauthorized creates a PcloudPvminstancesConsolePutUnauthorized with default headers values
func NewPcloudPvminstancesConsolePutUnauthorized() *PcloudPvminstancesConsolePutUnauthorized {
	return &PcloudPvminstancesConsolePutUnauthorized{}
}

/*
PcloudPvminstancesConsolePutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesConsolePutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console put unauthorized response has a 2xx status code
func (o *PcloudPvminstancesConsolePutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console put unauthorized response has a 3xx status code
func (o *PcloudPvminstancesConsolePutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console put unauthorized response has a 4xx status code
func (o *PcloudPvminstancesConsolePutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console put unauthorized response has a 5xx status code
func (o *PcloudPvminstancesConsolePutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console put unauthorized response a status code equal to that given
func (o *PcloudPvminstancesConsolePutUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudPvminstancesConsolePutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesConsolePutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesConsolePutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePutForbidden creates a PcloudPvminstancesConsolePutForbidden with default headers values
func NewPcloudPvminstancesConsolePutForbidden() *PcloudPvminstancesConsolePutForbidden {
	return &PcloudPvminstancesConsolePutForbidden{}
}

/*
PcloudPvminstancesConsolePutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesConsolePutForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console put forbidden response has a 2xx status code
func (o *PcloudPvminstancesConsolePutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console put forbidden response has a 3xx status code
func (o *PcloudPvminstancesConsolePutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console put forbidden response has a 4xx status code
func (o *PcloudPvminstancesConsolePutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console put forbidden response has a 5xx status code
func (o *PcloudPvminstancesConsolePutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console put forbidden response a status code equal to that given
func (o *PcloudPvminstancesConsolePutForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudPvminstancesConsolePutForbidden) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesConsolePutForbidden) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesConsolePutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePutNotFound creates a PcloudPvminstancesConsolePutNotFound with default headers values
func NewPcloudPvminstancesConsolePutNotFound() *PcloudPvminstancesConsolePutNotFound {
	return &PcloudPvminstancesConsolePutNotFound{}
}

/*
PcloudPvminstancesConsolePutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesConsolePutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console put not found response has a 2xx status code
func (o *PcloudPvminstancesConsolePutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console put not found response has a 3xx status code
func (o *PcloudPvminstancesConsolePutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console put not found response has a 4xx status code
func (o *PcloudPvminstancesConsolePutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances console put not found response has a 5xx status code
func (o *PcloudPvminstancesConsolePutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances console put not found response a status code equal to that given
func (o *PcloudPvminstancesConsolePutNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudPvminstancesConsolePutNotFound) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesConsolePutNotFound) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesConsolePutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePutInternalServerError creates a PcloudPvminstancesConsolePutInternalServerError with default headers values
func NewPcloudPvminstancesConsolePutInternalServerError() *PcloudPvminstancesConsolePutInternalServerError {
	return &PcloudPvminstancesConsolePutInternalServerError{}
}

/*
PcloudPvminstancesConsolePutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesConsolePutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances console put internal server error response has a 2xx status code
func (o *PcloudPvminstancesConsolePutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances console put internal server error response has a 3xx status code
func (o *PcloudPvminstancesConsolePutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances console put internal server error response has a 4xx status code
func (o *PcloudPvminstancesConsolePutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances console put internal server error response has a 5xx status code
func (o *PcloudPvminstancesConsolePutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances console put internal server error response a status code equal to that given
func (o *PcloudPvminstancesConsolePutInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudPvminstancesConsolePutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesConsolePutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesConsolePutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesConsolePutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

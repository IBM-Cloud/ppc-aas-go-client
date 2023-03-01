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

// PcloudPvminstancesVolumesSetbootPutReader is a Reader for the PcloudPvminstancesVolumesSetbootPut structure.
type PcloudPvminstancesVolumesSetbootPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesVolumesSetbootPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesVolumesSetbootPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesVolumesSetbootPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesVolumesSetbootPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesVolumesSetbootPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesVolumesSetbootPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesVolumesSetbootPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesVolumesSetbootPutOK creates a PcloudPvminstancesVolumesSetbootPutOK with default headers values
func NewPcloudPvminstancesVolumesSetbootPutOK() *PcloudPvminstancesVolumesSetbootPutOK {
	return &PcloudPvminstancesVolumesSetbootPutOK{}
}

/*
PcloudPvminstancesVolumesSetbootPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesVolumesSetbootPutOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud pvminstances volumes setboot put o k response has a 2xx status code
func (o *PcloudPvminstancesVolumesSetbootPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances volumes setboot put o k response has a 3xx status code
func (o *PcloudPvminstancesVolumesSetbootPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes setboot put o k response has a 4xx status code
func (o *PcloudPvminstancesVolumesSetbootPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances volumes setboot put o k response has a 5xx status code
func (o *PcloudPvminstancesVolumesSetbootPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes setboot put o k response a status code equal to that given
func (o *PcloudPvminstancesVolumesSetbootPutOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudPvminstancesVolumesSetbootPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutOK) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesSetbootPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutBadRequest creates a PcloudPvminstancesVolumesSetbootPutBadRequest with default headers values
func NewPcloudPvminstancesVolumesSetbootPutBadRequest() *PcloudPvminstancesVolumesSetbootPutBadRequest {
	return &PcloudPvminstancesVolumesSetbootPutBadRequest{}
}

/*
PcloudPvminstancesVolumesSetbootPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesVolumesSetbootPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes setboot put bad request response has a 2xx status code
func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes setboot put bad request response has a 3xx status code
func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes setboot put bad request response has a 4xx status code
func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes setboot put bad request response has a 5xx status code
func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes setboot put bad request response a status code equal to that given
func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutUnauthorized creates a PcloudPvminstancesVolumesSetbootPutUnauthorized with default headers values
func NewPcloudPvminstancesVolumesSetbootPutUnauthorized() *PcloudPvminstancesVolumesSetbootPutUnauthorized {
	return &PcloudPvminstancesVolumesSetbootPutUnauthorized{}
}

/*
PcloudPvminstancesVolumesSetbootPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesVolumesSetbootPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes setboot put unauthorized response has a 2xx status code
func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes setboot put unauthorized response has a 3xx status code
func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes setboot put unauthorized response has a 4xx status code
func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes setboot put unauthorized response has a 5xx status code
func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes setboot put unauthorized response a status code equal to that given
func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesSetbootPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutForbidden creates a PcloudPvminstancesVolumesSetbootPutForbidden with default headers values
func NewPcloudPvminstancesVolumesSetbootPutForbidden() *PcloudPvminstancesVolumesSetbootPutForbidden {
	return &PcloudPvminstancesVolumesSetbootPutForbidden{}
}

/*
PcloudPvminstancesVolumesSetbootPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesVolumesSetbootPutForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes setboot put forbidden response has a 2xx status code
func (o *PcloudPvminstancesVolumesSetbootPutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes setboot put forbidden response has a 3xx status code
func (o *PcloudPvminstancesVolumesSetbootPutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes setboot put forbidden response has a 4xx status code
func (o *PcloudPvminstancesVolumesSetbootPutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes setboot put forbidden response has a 5xx status code
func (o *PcloudPvminstancesVolumesSetbootPutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes setboot put forbidden response a status code equal to that given
func (o *PcloudPvminstancesVolumesSetbootPutForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudPvminstancesVolumesSetbootPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutForbidden) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesSetbootPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutNotFound creates a PcloudPvminstancesVolumesSetbootPutNotFound with default headers values
func NewPcloudPvminstancesVolumesSetbootPutNotFound() *PcloudPvminstancesVolumesSetbootPutNotFound {
	return &PcloudPvminstancesVolumesSetbootPutNotFound{}
}

/*
PcloudPvminstancesVolumesSetbootPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesVolumesSetbootPutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes setboot put not found response has a 2xx status code
func (o *PcloudPvminstancesVolumesSetbootPutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes setboot put not found response has a 3xx status code
func (o *PcloudPvminstancesVolumesSetbootPutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes setboot put not found response has a 4xx status code
func (o *PcloudPvminstancesVolumesSetbootPutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes setboot put not found response has a 5xx status code
func (o *PcloudPvminstancesVolumesSetbootPutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes setboot put not found response a status code equal to that given
func (o *PcloudPvminstancesVolumesSetbootPutNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudPvminstancesVolumesSetbootPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutNotFound) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesSetbootPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutInternalServerError creates a PcloudPvminstancesVolumesSetbootPutInternalServerError with default headers values
func NewPcloudPvminstancesVolumesSetbootPutInternalServerError() *PcloudPvminstancesVolumesSetbootPutInternalServerError {
	return &PcloudPvminstancesVolumesSetbootPutInternalServerError{}
}

/*
PcloudPvminstancesVolumesSetbootPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesVolumesSetbootPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes setboot put internal server error response has a 2xx status code
func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes setboot put internal server error response has a 3xx status code
func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes setboot put internal server error response has a 4xx status code
func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances volumes setboot put internal server error response has a 5xx status code
func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances volumes setboot put internal server error response a status code equal to that given
func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudCloudinstancesPutReader is a Reader for the PcloudCloudinstancesPut structure.
type PcloudCloudinstancesPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudCloudinstancesPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesPutOK creates a PcloudCloudinstancesPutOK with default headers values
func NewPcloudCloudinstancesPutOK() *PcloudCloudinstancesPutOK {
	return &PcloudCloudinstancesPutOK{}
}

/*
PcloudCloudinstancesPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesPutOK struct {
	Payload *models.CloudInstance
}

// IsSuccess returns true when this pcloud cloudinstances put o k response has a 2xx status code
func (o *PcloudCloudinstancesPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances put o k response has a 3xx status code
func (o *PcloudCloudinstancesPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances put o k response has a 4xx status code
func (o *PcloudCloudinstancesPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances put o k response has a 5xx status code
func (o *PcloudCloudinstancesPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances put o k response a status code equal to that given
func (o *PcloudCloudinstancesPutOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudCloudinstancesPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesPutOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesPutOK) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesPutOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesPutOK) GetPayload() *models.CloudInstance {
	return o.Payload
}

func (o *PcloudCloudinstancesPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesPutBadRequest creates a PcloudCloudinstancesPutBadRequest with default headers values
func NewPcloudCloudinstancesPutBadRequest() *PcloudCloudinstancesPutBadRequest {
	return &PcloudCloudinstancesPutBadRequest{}
}

/*
PcloudCloudinstancesPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances put bad request response has a 2xx status code
func (o *PcloudCloudinstancesPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances put bad request response has a 3xx status code
func (o *PcloudCloudinstancesPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances put bad request response has a 4xx status code
func (o *PcloudCloudinstancesPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances put bad request response has a 5xx status code
func (o *PcloudCloudinstancesPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances put bad request response a status code equal to that given
func (o *PcloudCloudinstancesPutBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudCloudinstancesPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesPutUnauthorized creates a PcloudCloudinstancesPutUnauthorized with default headers values
func NewPcloudCloudinstancesPutUnauthorized() *PcloudCloudinstancesPutUnauthorized {
	return &PcloudCloudinstancesPutUnauthorized{}
}

/*
PcloudCloudinstancesPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances put unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances put unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances put unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances put unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances put unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudCloudinstancesPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesPutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesPutUnprocessableEntity creates a PcloudCloudinstancesPutUnprocessableEntity with default headers values
func NewPcloudCloudinstancesPutUnprocessableEntity() *PcloudCloudinstancesPutUnprocessableEntity {
	return &PcloudCloudinstancesPutUnprocessableEntity{}
}

/*
PcloudCloudinstancesPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudCloudinstancesPutUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances put unprocessable entity response has a 2xx status code
func (o *PcloudCloudinstancesPutUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances put unprocessable entity response has a 3xx status code
func (o *PcloudCloudinstancesPutUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances put unprocessable entity response has a 4xx status code
func (o *PcloudCloudinstancesPutUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances put unprocessable entity response has a 5xx status code
func (o *PcloudCloudinstancesPutUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances put unprocessable entity response a status code equal to that given
func (o *PcloudCloudinstancesPutUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PcloudCloudinstancesPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudCloudinstancesPutUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudCloudinstancesPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesPutInternalServerError creates a PcloudCloudinstancesPutInternalServerError with default headers values
func NewPcloudCloudinstancesPutInternalServerError() *PcloudCloudinstancesPutInternalServerError {
	return &PcloudCloudinstancesPutInternalServerError{}
}

/*
PcloudCloudinstancesPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances put internal server error response has a 2xx status code
func (o *PcloudCloudinstancesPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances put internal server error response has a 3xx status code
func (o *PcloudCloudinstancesPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances put internal server error response has a 4xx status code
func (o *PcloudCloudinstancesPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances put internal server error response has a 5xx status code
func (o *PcloudCloudinstancesPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances put internal server error response a status code equal to that given
func (o *PcloudCloudinstancesPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudCloudinstancesPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesPutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

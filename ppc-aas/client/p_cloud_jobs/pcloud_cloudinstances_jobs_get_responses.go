// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudCloudinstancesJobsGetReader is a Reader for the PcloudCloudinstancesJobsGet structure.
type PcloudCloudinstancesJobsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesJobsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesJobsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesJobsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesJobsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesJobsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesJobsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesJobsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesJobsGetOK creates a PcloudCloudinstancesJobsGetOK with default headers values
func NewPcloudCloudinstancesJobsGetOK() *PcloudCloudinstancesJobsGetOK {
	return &PcloudCloudinstancesJobsGetOK{}
}

/*
PcloudCloudinstancesJobsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesJobsGetOK struct {
	Payload *models.Job
}

// IsSuccess returns true when this pcloud cloudinstances jobs get o k response has a 2xx status code
func (o *PcloudCloudinstancesJobsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances jobs get o k response has a 3xx status code
func (o *PcloudCloudinstancesJobsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs get o k response has a 4xx status code
func (o *PcloudCloudinstancesJobsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances jobs get o k response has a 5xx status code
func (o *PcloudCloudinstancesJobsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs get o k response a status code equal to that given
func (o *PcloudCloudinstancesJobsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud cloudinstances jobs get o k response
func (o *PcloudCloudinstancesJobsGetOK) Code() int {
	return 200
}

func (o *PcloudCloudinstancesJobsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetOK) GetPayload() *models.Job {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetBadRequest creates a PcloudCloudinstancesJobsGetBadRequest with default headers values
func NewPcloudCloudinstancesJobsGetBadRequest() *PcloudCloudinstancesJobsGetBadRequest {
	return &PcloudCloudinstancesJobsGetBadRequest{}
}

/*
PcloudCloudinstancesJobsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesJobsGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs get bad request response has a 2xx status code
func (o *PcloudCloudinstancesJobsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs get bad request response has a 3xx status code
func (o *PcloudCloudinstancesJobsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs get bad request response has a 4xx status code
func (o *PcloudCloudinstancesJobsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances jobs get bad request response has a 5xx status code
func (o *PcloudCloudinstancesJobsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs get bad request response a status code equal to that given
func (o *PcloudCloudinstancesJobsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudinstances jobs get bad request response
func (o *PcloudCloudinstancesJobsGetBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudinstancesJobsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetUnauthorized creates a PcloudCloudinstancesJobsGetUnauthorized with default headers values
func NewPcloudCloudinstancesJobsGetUnauthorized() *PcloudCloudinstancesJobsGetUnauthorized {
	return &PcloudCloudinstancesJobsGetUnauthorized{}
}

/*
PcloudCloudinstancesJobsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesJobsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs get unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesJobsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs get unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesJobsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs get unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesJobsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances jobs get unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesJobsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs get unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesJobsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudinstances jobs get unauthorized response
func (o *PcloudCloudinstancesJobsGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudinstancesJobsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetForbidden creates a PcloudCloudinstancesJobsGetForbidden with default headers values
func NewPcloudCloudinstancesJobsGetForbidden() *PcloudCloudinstancesJobsGetForbidden {
	return &PcloudCloudinstancesJobsGetForbidden{}
}

/*
PcloudCloudinstancesJobsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesJobsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs get forbidden response has a 2xx status code
func (o *PcloudCloudinstancesJobsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs get forbidden response has a 3xx status code
func (o *PcloudCloudinstancesJobsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs get forbidden response has a 4xx status code
func (o *PcloudCloudinstancesJobsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances jobs get forbidden response has a 5xx status code
func (o *PcloudCloudinstancesJobsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs get forbidden response a status code equal to that given
func (o *PcloudCloudinstancesJobsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud cloudinstances jobs get forbidden response
func (o *PcloudCloudinstancesJobsGetForbidden) Code() int {
	return 403
}

func (o *PcloudCloudinstancesJobsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetNotFound creates a PcloudCloudinstancesJobsGetNotFound with default headers values
func NewPcloudCloudinstancesJobsGetNotFound() *PcloudCloudinstancesJobsGetNotFound {
	return &PcloudCloudinstancesJobsGetNotFound{}
}

/*
PcloudCloudinstancesJobsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesJobsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs get not found response has a 2xx status code
func (o *PcloudCloudinstancesJobsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs get not found response has a 3xx status code
func (o *PcloudCloudinstancesJobsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs get not found response has a 4xx status code
func (o *PcloudCloudinstancesJobsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances jobs get not found response has a 5xx status code
func (o *PcloudCloudinstancesJobsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs get not found response a status code equal to that given
func (o *PcloudCloudinstancesJobsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud cloudinstances jobs get not found response
func (o *PcloudCloudinstancesJobsGetNotFound) Code() int {
	return 404
}

func (o *PcloudCloudinstancesJobsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetInternalServerError creates a PcloudCloudinstancesJobsGetInternalServerError with default headers values
func NewPcloudCloudinstancesJobsGetInternalServerError() *PcloudCloudinstancesJobsGetInternalServerError {
	return &PcloudCloudinstancesJobsGetInternalServerError{}
}

/*
PcloudCloudinstancesJobsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesJobsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs get internal server error response has a 2xx status code
func (o *PcloudCloudinstancesJobsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs get internal server error response has a 3xx status code
func (o *PcloudCloudinstancesJobsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs get internal server error response has a 4xx status code
func (o *PcloudCloudinstancesJobsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances jobs get internal server error response has a 5xx status code
func (o *PcloudCloudinstancesJobsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances jobs get internal server error response a status code equal to that given
func (o *PcloudCloudinstancesJobsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudinstances jobs get internal server error response
func (o *PcloudCloudinstancesJobsGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudinstancesJobsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

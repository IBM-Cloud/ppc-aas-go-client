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

// PcloudCloudinstancesDeleteReader is a Reader for the PcloudCloudinstancesDelete structure.
type PcloudCloudinstancesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudCloudinstancesDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesDeleteOK creates a PcloudCloudinstancesDeleteOK with default headers values
func NewPcloudCloudinstancesDeleteOK() *PcloudCloudinstancesDeleteOK {
	return &PcloudCloudinstancesDeleteOK{}
}

/*
PcloudCloudinstancesDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesDeleteOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud cloudinstances delete o k response has a 2xx status code
func (o *PcloudCloudinstancesDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances delete o k response has a 3xx status code
func (o *PcloudCloudinstancesDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances delete o k response has a 4xx status code
func (o *PcloudCloudinstancesDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances delete o k response has a 5xx status code
func (o *PcloudCloudinstancesDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances delete o k response a status code equal to that given
func (o *PcloudCloudinstancesDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud cloudinstances delete o k response
func (o *PcloudCloudinstancesDeleteOK) Code() int {
	return 200
}

func (o *PcloudCloudinstancesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudCloudinstancesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesDeleteBadRequest creates a PcloudCloudinstancesDeleteBadRequest with default headers values
func NewPcloudCloudinstancesDeleteBadRequest() *PcloudCloudinstancesDeleteBadRequest {
	return &PcloudCloudinstancesDeleteBadRequest{}
}

/*
PcloudCloudinstancesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances delete bad request response has a 2xx status code
func (o *PcloudCloudinstancesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances delete bad request response has a 3xx status code
func (o *PcloudCloudinstancesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances delete bad request response has a 4xx status code
func (o *PcloudCloudinstancesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances delete bad request response has a 5xx status code
func (o *PcloudCloudinstancesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances delete bad request response a status code equal to that given
func (o *PcloudCloudinstancesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudinstances delete bad request response
func (o *PcloudCloudinstancesDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudinstancesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesDeleteUnauthorized creates a PcloudCloudinstancesDeleteUnauthorized with default headers values
func NewPcloudCloudinstancesDeleteUnauthorized() *PcloudCloudinstancesDeleteUnauthorized {
	return &PcloudCloudinstancesDeleteUnauthorized{}
}

/*
PcloudCloudinstancesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances delete unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances delete unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances delete unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances delete unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances delete unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudinstances delete unauthorized response
func (o *PcloudCloudinstancesDeleteUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudinstancesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesDeleteGone creates a PcloudCloudinstancesDeleteGone with default headers values
func NewPcloudCloudinstancesDeleteGone() *PcloudCloudinstancesDeleteGone {
	return &PcloudCloudinstancesDeleteGone{}
}

/*
PcloudCloudinstancesDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudCloudinstancesDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances delete gone response has a 2xx status code
func (o *PcloudCloudinstancesDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances delete gone response has a 3xx status code
func (o *PcloudCloudinstancesDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances delete gone response has a 4xx status code
func (o *PcloudCloudinstancesDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances delete gone response has a 5xx status code
func (o *PcloudCloudinstancesDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances delete gone response a status code equal to that given
func (o *PcloudCloudinstancesDeleteGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the pcloud cloudinstances delete gone response
func (o *PcloudCloudinstancesDeleteGone) Code() int {
	return 410
}

func (o *PcloudCloudinstancesDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudCloudinstancesDeleteGone) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudCloudinstancesDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesDeleteInternalServerError creates a PcloudCloudinstancesDeleteInternalServerError with default headers values
func NewPcloudCloudinstancesDeleteInternalServerError() *PcloudCloudinstancesDeleteInternalServerError {
	return &PcloudCloudinstancesDeleteInternalServerError{}
}

/*
PcloudCloudinstancesDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances delete internal server error response has a 2xx status code
func (o *PcloudCloudinstancesDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances delete internal server error response has a 3xx status code
func (o *PcloudCloudinstancesDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances delete internal server error response has a 4xx status code
func (o *PcloudCloudinstancesDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances delete internal server error response has a 5xx status code
func (o *PcloudCloudinstancesDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances delete internal server error response a status code equal to that given
func (o *PcloudCloudinstancesDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudinstances delete internal server error response
func (o *PcloudCloudinstancesDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudinstancesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

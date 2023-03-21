// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudCloudinstancesSnapshotsDeleteReader is a Reader for the PcloudCloudinstancesSnapshotsDelete structure.
type PcloudCloudinstancesSnapshotsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesSnapshotsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudCloudinstancesSnapshotsDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesSnapshotsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesSnapshotsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesSnapshotsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesSnapshotsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudCloudinstancesSnapshotsDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesSnapshotsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesSnapshotsDeleteAccepted creates a PcloudCloudinstancesSnapshotsDeleteAccepted with default headers values
func NewPcloudCloudinstancesSnapshotsDeleteAccepted() *PcloudCloudinstancesSnapshotsDeleteAccepted {
	return &PcloudCloudinstancesSnapshotsDeleteAccepted{}
}

/*
PcloudCloudinstancesSnapshotsDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudCloudinstancesSnapshotsDeleteAccepted struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud cloudinstances snapshots delete accepted response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances snapshots delete accepted response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots delete accepted response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances snapshots delete accepted response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots delete accepted response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud cloudinstances snapshots delete accepted response
func (o *PcloudCloudinstancesSnapshotsDeleteAccepted) Code() int {
	return 202
}

func (o *PcloudCloudinstancesSnapshotsDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteAccepted  %+v", 202, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteAccepted) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteAccepted  %+v", 202, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteAccepted) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsDeleteBadRequest creates a PcloudCloudinstancesSnapshotsDeleteBadRequest with default headers values
func NewPcloudCloudinstancesSnapshotsDeleteBadRequest() *PcloudCloudinstancesSnapshotsDeleteBadRequest {
	return &PcloudCloudinstancesSnapshotsDeleteBadRequest{}
}

/*
PcloudCloudinstancesSnapshotsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesSnapshotsDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots delete bad request response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots delete bad request response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots delete bad request response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots delete bad request response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots delete bad request response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudinstances snapshots delete bad request response
func (o *PcloudCloudinstancesSnapshotsDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudinstancesSnapshotsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsDeleteUnauthorized creates a PcloudCloudinstancesSnapshotsDeleteUnauthorized with default headers values
func NewPcloudCloudinstancesSnapshotsDeleteUnauthorized() *PcloudCloudinstancesSnapshotsDeleteUnauthorized {
	return &PcloudCloudinstancesSnapshotsDeleteUnauthorized{}
}

/*
PcloudCloudinstancesSnapshotsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesSnapshotsDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots delete unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots delete unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots delete unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots delete unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots delete unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudinstances snapshots delete unauthorized response
func (o *PcloudCloudinstancesSnapshotsDeleteUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudinstancesSnapshotsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsDeleteForbidden creates a PcloudCloudinstancesSnapshotsDeleteForbidden with default headers values
func NewPcloudCloudinstancesSnapshotsDeleteForbidden() *PcloudCloudinstancesSnapshotsDeleteForbidden {
	return &PcloudCloudinstancesSnapshotsDeleteForbidden{}
}

/*
PcloudCloudinstancesSnapshotsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesSnapshotsDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots delete forbidden response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots delete forbidden response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots delete forbidden response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots delete forbidden response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots delete forbidden response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud cloudinstances snapshots delete forbidden response
func (o *PcloudCloudinstancesSnapshotsDeleteForbidden) Code() int {
	return 403
}

func (o *PcloudCloudinstancesSnapshotsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsDeleteNotFound creates a PcloudCloudinstancesSnapshotsDeleteNotFound with default headers values
func NewPcloudCloudinstancesSnapshotsDeleteNotFound() *PcloudCloudinstancesSnapshotsDeleteNotFound {
	return &PcloudCloudinstancesSnapshotsDeleteNotFound{}
}

/*
PcloudCloudinstancesSnapshotsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesSnapshotsDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots delete not found response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots delete not found response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots delete not found response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots delete not found response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots delete not found response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud cloudinstances snapshots delete not found response
func (o *PcloudCloudinstancesSnapshotsDeleteNotFound) Code() int {
	return 404
}

func (o *PcloudCloudinstancesSnapshotsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsDeleteGone creates a PcloudCloudinstancesSnapshotsDeleteGone with default headers values
func NewPcloudCloudinstancesSnapshotsDeleteGone() *PcloudCloudinstancesSnapshotsDeleteGone {
	return &PcloudCloudinstancesSnapshotsDeleteGone{}
}

/*
PcloudCloudinstancesSnapshotsDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudCloudinstancesSnapshotsDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots delete gone response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots delete gone response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots delete gone response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots delete gone response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots delete gone response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsDeleteGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the pcloud cloudinstances snapshots delete gone response
func (o *PcloudCloudinstancesSnapshotsDeleteGone) Code() int {
	return 410
}

func (o *PcloudCloudinstancesSnapshotsDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteGone) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsDeleteInternalServerError creates a PcloudCloudinstancesSnapshotsDeleteInternalServerError with default headers values
func NewPcloudCloudinstancesSnapshotsDeleteInternalServerError() *PcloudCloudinstancesSnapshotsDeleteInternalServerError {
	return &PcloudCloudinstancesSnapshotsDeleteInternalServerError{}
}

/*
PcloudCloudinstancesSnapshotsDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesSnapshotsDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots delete internal server error response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots delete internal server error response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots delete internal server error response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances snapshots delete internal server error response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances snapshots delete internal server error response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudinstances snapshots delete internal server error response
func (o *PcloudCloudinstancesSnapshotsDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudinstancesSnapshotsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

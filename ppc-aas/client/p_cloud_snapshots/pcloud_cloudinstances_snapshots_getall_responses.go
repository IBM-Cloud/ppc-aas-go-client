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

// PcloudCloudinstancesSnapshotsGetallReader is a Reader for the PcloudCloudinstancesSnapshotsGetall structure.
type PcloudCloudinstancesSnapshotsGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesSnapshotsGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesSnapshotsGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesSnapshotsGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesSnapshotsGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesSnapshotsGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesSnapshotsGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesSnapshotsGetallOK creates a PcloudCloudinstancesSnapshotsGetallOK with default headers values
func NewPcloudCloudinstancesSnapshotsGetallOK() *PcloudCloudinstancesSnapshotsGetallOK {
	return &PcloudCloudinstancesSnapshotsGetallOK{}
}

/*
PcloudCloudinstancesSnapshotsGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesSnapshotsGetallOK struct {
	Payload *models.Snapshots
}

// IsSuccess returns true when this pcloud cloudinstances snapshots getall o k response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances snapshots getall o k response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots getall o k response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances snapshots getall o k response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots getall o k response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud cloudinstances snapshots getall o k response
func (o *PcloudCloudinstancesSnapshotsGetallOK) Code() int {
	return 200
}

func (o *PcloudCloudinstancesSnapshotsGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetallOK) GetPayload() *models.Snapshots {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Snapshots)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsGetallBadRequest creates a PcloudCloudinstancesSnapshotsGetallBadRequest with default headers values
func NewPcloudCloudinstancesSnapshotsGetallBadRequest() *PcloudCloudinstancesSnapshotsGetallBadRequest {
	return &PcloudCloudinstancesSnapshotsGetallBadRequest{}
}

/*
PcloudCloudinstancesSnapshotsGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesSnapshotsGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots getall bad request response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots getall bad request response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots getall bad request response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots getall bad request response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots getall bad request response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudinstances snapshots getall bad request response
func (o *PcloudCloudinstancesSnapshotsGetallBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudinstancesSnapshotsGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsGetallUnauthorized creates a PcloudCloudinstancesSnapshotsGetallUnauthorized with default headers values
func NewPcloudCloudinstancesSnapshotsGetallUnauthorized() *PcloudCloudinstancesSnapshotsGetallUnauthorized {
	return &PcloudCloudinstancesSnapshotsGetallUnauthorized{}
}

/*
PcloudCloudinstancesSnapshotsGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesSnapshotsGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots getall unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots getall unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots getall unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots getall unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots getall unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudinstances snapshots getall unauthorized response
func (o *PcloudCloudinstancesSnapshotsGetallUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudinstancesSnapshotsGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsGetallForbidden creates a PcloudCloudinstancesSnapshotsGetallForbidden with default headers values
func NewPcloudCloudinstancesSnapshotsGetallForbidden() *PcloudCloudinstancesSnapshotsGetallForbidden {
	return &PcloudCloudinstancesSnapshotsGetallForbidden{}
}

/*
PcloudCloudinstancesSnapshotsGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesSnapshotsGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots getall forbidden response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots getall forbidden response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots getall forbidden response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots getall forbidden response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots getall forbidden response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsGetallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud cloudinstances snapshots getall forbidden response
func (o *PcloudCloudinstancesSnapshotsGetallForbidden) Code() int {
	return 403
}

func (o *PcloudCloudinstancesSnapshotsGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsGetallInternalServerError creates a PcloudCloudinstancesSnapshotsGetallInternalServerError with default headers values
func NewPcloudCloudinstancesSnapshotsGetallInternalServerError() *PcloudCloudinstancesSnapshotsGetallInternalServerError {
	return &PcloudCloudinstancesSnapshotsGetallInternalServerError{}
}

/*
PcloudCloudinstancesSnapshotsGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesSnapshotsGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots getall internal server error response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots getall internal server error response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots getall internal server error response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances snapshots getall internal server error response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances snapshots getall internal server error response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudinstances snapshots getall internal server error response
func (o *PcloudCloudinstancesSnapshotsGetallInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudinstancesSnapshotsGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots][%d] pcloudCloudinstancesSnapshotsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_storage_capacity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudStoragecapacityPoolsGetallReader is a Reader for the PcloudStoragecapacityPoolsGetall structure.
type PcloudStoragecapacityPoolsGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudStoragecapacityPoolsGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudStoragecapacityPoolsGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPcloudStoragecapacityPoolsGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudStoragecapacityPoolsGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudStoragecapacityPoolsGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudStoragecapacityPoolsGetallOK creates a PcloudStoragecapacityPoolsGetallOK with default headers values
func NewPcloudStoragecapacityPoolsGetallOK() *PcloudStoragecapacityPoolsGetallOK {
	return &PcloudStoragecapacityPoolsGetallOK{}
}

/*
PcloudStoragecapacityPoolsGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudStoragecapacityPoolsGetallOK struct {
	Payload *models.StoragePoolsCapacity
}

// IsSuccess returns true when this pcloud storagecapacity pools getall o k response has a 2xx status code
func (o *PcloudStoragecapacityPoolsGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud storagecapacity pools getall o k response has a 3xx status code
func (o *PcloudStoragecapacityPoolsGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity pools getall o k response has a 4xx status code
func (o *PcloudStoragecapacityPoolsGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud storagecapacity pools getall o k response has a 5xx status code
func (o *PcloudStoragecapacityPoolsGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud storagecapacity pools getall o k response a status code equal to that given
func (o *PcloudStoragecapacityPoolsGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud storagecapacity pools getall o k response
func (o *PcloudStoragecapacityPoolsGetallOK) Code() int {
	return 200
}

func (o *PcloudStoragecapacityPoolsGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools][%d] pcloudStoragecapacityPoolsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudStoragecapacityPoolsGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools][%d] pcloudStoragecapacityPoolsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudStoragecapacityPoolsGetallOK) GetPayload() *models.StoragePoolsCapacity {
	return o.Payload
}

func (o *PcloudStoragecapacityPoolsGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePoolsCapacity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityPoolsGetallUnauthorized creates a PcloudStoragecapacityPoolsGetallUnauthorized with default headers values
func NewPcloudStoragecapacityPoolsGetallUnauthorized() *PcloudStoragecapacityPoolsGetallUnauthorized {
	return &PcloudStoragecapacityPoolsGetallUnauthorized{}
}

/*
PcloudStoragecapacityPoolsGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudStoragecapacityPoolsGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud storagecapacity pools getall unauthorized response has a 2xx status code
func (o *PcloudStoragecapacityPoolsGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud storagecapacity pools getall unauthorized response has a 3xx status code
func (o *PcloudStoragecapacityPoolsGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity pools getall unauthorized response has a 4xx status code
func (o *PcloudStoragecapacityPoolsGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud storagecapacity pools getall unauthorized response has a 5xx status code
func (o *PcloudStoragecapacityPoolsGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud storagecapacity pools getall unauthorized response a status code equal to that given
func (o *PcloudStoragecapacityPoolsGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud storagecapacity pools getall unauthorized response
func (o *PcloudStoragecapacityPoolsGetallUnauthorized) Code() int {
	return 401
}

func (o *PcloudStoragecapacityPoolsGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools][%d] pcloudStoragecapacityPoolsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudStoragecapacityPoolsGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools][%d] pcloudStoragecapacityPoolsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudStoragecapacityPoolsGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudStoragecapacityPoolsGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityPoolsGetallForbidden creates a PcloudStoragecapacityPoolsGetallForbidden with default headers values
func NewPcloudStoragecapacityPoolsGetallForbidden() *PcloudStoragecapacityPoolsGetallForbidden {
	return &PcloudStoragecapacityPoolsGetallForbidden{}
}

/*
PcloudStoragecapacityPoolsGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudStoragecapacityPoolsGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud storagecapacity pools getall forbidden response has a 2xx status code
func (o *PcloudStoragecapacityPoolsGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud storagecapacity pools getall forbidden response has a 3xx status code
func (o *PcloudStoragecapacityPoolsGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity pools getall forbidden response has a 4xx status code
func (o *PcloudStoragecapacityPoolsGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud storagecapacity pools getall forbidden response has a 5xx status code
func (o *PcloudStoragecapacityPoolsGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud storagecapacity pools getall forbidden response a status code equal to that given
func (o *PcloudStoragecapacityPoolsGetallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud storagecapacity pools getall forbidden response
func (o *PcloudStoragecapacityPoolsGetallForbidden) Code() int {
	return 403
}

func (o *PcloudStoragecapacityPoolsGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools][%d] pcloudStoragecapacityPoolsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudStoragecapacityPoolsGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools][%d] pcloudStoragecapacityPoolsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudStoragecapacityPoolsGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudStoragecapacityPoolsGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityPoolsGetallInternalServerError creates a PcloudStoragecapacityPoolsGetallInternalServerError with default headers values
func NewPcloudStoragecapacityPoolsGetallInternalServerError() *PcloudStoragecapacityPoolsGetallInternalServerError {
	return &PcloudStoragecapacityPoolsGetallInternalServerError{}
}

/*
PcloudStoragecapacityPoolsGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudStoragecapacityPoolsGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud storagecapacity pools getall internal server error response has a 2xx status code
func (o *PcloudStoragecapacityPoolsGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud storagecapacity pools getall internal server error response has a 3xx status code
func (o *PcloudStoragecapacityPoolsGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity pools getall internal server error response has a 4xx status code
func (o *PcloudStoragecapacityPoolsGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud storagecapacity pools getall internal server error response has a 5xx status code
func (o *PcloudStoragecapacityPoolsGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud storagecapacity pools getall internal server error response a status code equal to that given
func (o *PcloudStoragecapacityPoolsGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud storagecapacity pools getall internal server error response
func (o *PcloudStoragecapacityPoolsGetallInternalServerError) Code() int {
	return 500
}

func (o *PcloudStoragecapacityPoolsGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools][%d] pcloudStoragecapacityPoolsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudStoragecapacityPoolsGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools][%d] pcloudStoragecapacityPoolsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudStoragecapacityPoolsGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudStoragecapacityPoolsGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

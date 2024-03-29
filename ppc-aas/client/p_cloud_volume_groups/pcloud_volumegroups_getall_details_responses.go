// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volume_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudVolumegroupsGetallDetailsReader is a Reader for the PcloudVolumegroupsGetallDetails structure.
type PcloudVolumegroupsGetallDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVolumegroupsGetallDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVolumegroupsGetallDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVolumegroupsGetallDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVolumegroupsGetallDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVolumegroupsGetallDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVolumegroupsGetallDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/details] pcloud.volumegroups.getallDetails", response, response.Code())
	}
}

// NewPcloudVolumegroupsGetallDetailsOK creates a PcloudVolumegroupsGetallDetailsOK with default headers values
func NewPcloudVolumegroupsGetallDetailsOK() *PcloudVolumegroupsGetallDetailsOK {
	return &PcloudVolumegroupsGetallDetailsOK{}
}

/*
PcloudVolumegroupsGetallDetailsOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVolumegroupsGetallDetailsOK struct {
	Payload *models.VolumeGroupsDetails
}

// IsSuccess returns true when this pcloud volumegroups getall details o k response has a 2xx status code
func (o *PcloudVolumegroupsGetallDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud volumegroups getall details o k response has a 3xx status code
func (o *PcloudVolumegroupsGetallDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups getall details o k response has a 4xx status code
func (o *PcloudVolumegroupsGetallDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volumegroups getall details o k response has a 5xx status code
func (o *PcloudVolumegroupsGetallDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups getall details o k response a status code equal to that given
func (o *PcloudVolumegroupsGetallDetailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud volumegroups getall details o k response
func (o *PcloudVolumegroupsGetallDetailsOK) Code() int {
	return 200
}

func (o *PcloudVolumegroupsGetallDetailsOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/details][%d] pcloudVolumegroupsGetallDetailsOK  %+v", 200, o.Payload)
}

func (o *PcloudVolumegroupsGetallDetailsOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/details][%d] pcloudVolumegroupsGetallDetailsOK  %+v", 200, o.Payload)
}

func (o *PcloudVolumegroupsGetallDetailsOK) GetPayload() *models.VolumeGroupsDetails {
	return o.Payload
}

func (o *PcloudVolumegroupsGetallDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupsDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetallDetailsBadRequest creates a PcloudVolumegroupsGetallDetailsBadRequest with default headers values
func NewPcloudVolumegroupsGetallDetailsBadRequest() *PcloudVolumegroupsGetallDetailsBadRequest {
	return &PcloudVolumegroupsGetallDetailsBadRequest{}
}

/*
PcloudVolumegroupsGetallDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVolumegroupsGetallDetailsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups getall details bad request response has a 2xx status code
func (o *PcloudVolumegroupsGetallDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups getall details bad request response has a 3xx status code
func (o *PcloudVolumegroupsGetallDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups getall details bad request response has a 4xx status code
func (o *PcloudVolumegroupsGetallDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups getall details bad request response has a 5xx status code
func (o *PcloudVolumegroupsGetallDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups getall details bad request response a status code equal to that given
func (o *PcloudVolumegroupsGetallDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud volumegroups getall details bad request response
func (o *PcloudVolumegroupsGetallDetailsBadRequest) Code() int {
	return 400
}

func (o *PcloudVolumegroupsGetallDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/details][%d] pcloudVolumegroupsGetallDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVolumegroupsGetallDetailsBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/details][%d] pcloudVolumegroupsGetallDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVolumegroupsGetallDetailsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetallDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetallDetailsUnauthorized creates a PcloudVolumegroupsGetallDetailsUnauthorized with default headers values
func NewPcloudVolumegroupsGetallDetailsUnauthorized() *PcloudVolumegroupsGetallDetailsUnauthorized {
	return &PcloudVolumegroupsGetallDetailsUnauthorized{}
}

/*
PcloudVolumegroupsGetallDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVolumegroupsGetallDetailsUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups getall details unauthorized response has a 2xx status code
func (o *PcloudVolumegroupsGetallDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups getall details unauthorized response has a 3xx status code
func (o *PcloudVolumegroupsGetallDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups getall details unauthorized response has a 4xx status code
func (o *PcloudVolumegroupsGetallDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups getall details unauthorized response has a 5xx status code
func (o *PcloudVolumegroupsGetallDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups getall details unauthorized response a status code equal to that given
func (o *PcloudVolumegroupsGetallDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud volumegroups getall details unauthorized response
func (o *PcloudVolumegroupsGetallDetailsUnauthorized) Code() int {
	return 401
}

func (o *PcloudVolumegroupsGetallDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/details][%d] pcloudVolumegroupsGetallDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVolumegroupsGetallDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/details][%d] pcloudVolumegroupsGetallDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVolumegroupsGetallDetailsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetallDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetallDetailsForbidden creates a PcloudVolumegroupsGetallDetailsForbidden with default headers values
func NewPcloudVolumegroupsGetallDetailsForbidden() *PcloudVolumegroupsGetallDetailsForbidden {
	return &PcloudVolumegroupsGetallDetailsForbidden{}
}

/*
PcloudVolumegroupsGetallDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVolumegroupsGetallDetailsForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups getall details forbidden response has a 2xx status code
func (o *PcloudVolumegroupsGetallDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups getall details forbidden response has a 3xx status code
func (o *PcloudVolumegroupsGetallDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups getall details forbidden response has a 4xx status code
func (o *PcloudVolumegroupsGetallDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups getall details forbidden response has a 5xx status code
func (o *PcloudVolumegroupsGetallDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups getall details forbidden response a status code equal to that given
func (o *PcloudVolumegroupsGetallDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud volumegroups getall details forbidden response
func (o *PcloudVolumegroupsGetallDetailsForbidden) Code() int {
	return 403
}

func (o *PcloudVolumegroupsGetallDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/details][%d] pcloudVolumegroupsGetallDetailsForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVolumegroupsGetallDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/details][%d] pcloudVolumegroupsGetallDetailsForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVolumegroupsGetallDetailsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetallDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetallDetailsInternalServerError creates a PcloudVolumegroupsGetallDetailsInternalServerError with default headers values
func NewPcloudVolumegroupsGetallDetailsInternalServerError() *PcloudVolumegroupsGetallDetailsInternalServerError {
	return &PcloudVolumegroupsGetallDetailsInternalServerError{}
}

/*
PcloudVolumegroupsGetallDetailsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVolumegroupsGetallDetailsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups getall details internal server error response has a 2xx status code
func (o *PcloudVolumegroupsGetallDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups getall details internal server error response has a 3xx status code
func (o *PcloudVolumegroupsGetallDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups getall details internal server error response has a 4xx status code
func (o *PcloudVolumegroupsGetallDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volumegroups getall details internal server error response has a 5xx status code
func (o *PcloudVolumegroupsGetallDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud volumegroups getall details internal server error response a status code equal to that given
func (o *PcloudVolumegroupsGetallDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud volumegroups getall details internal server error response
func (o *PcloudVolumegroupsGetallDetailsInternalServerError) Code() int {
	return 500
}

func (o *PcloudVolumegroupsGetallDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/details][%d] pcloudVolumegroupsGetallDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVolumegroupsGetallDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/details][%d] pcloudVolumegroupsGetallDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVolumegroupsGetallDetailsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetallDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

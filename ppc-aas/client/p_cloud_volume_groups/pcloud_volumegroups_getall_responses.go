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

// PcloudVolumegroupsGetallReader is a Reader for the PcloudVolumegroupsGetall structure.
type PcloudVolumegroupsGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVolumegroupsGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVolumegroupsGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVolumegroupsGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVolumegroupsGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVolumegroupsGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVolumegroupsGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudVolumegroupsGetallOK creates a PcloudVolumegroupsGetallOK with default headers values
func NewPcloudVolumegroupsGetallOK() *PcloudVolumegroupsGetallOK {
	return &PcloudVolumegroupsGetallOK{}
}

/*
PcloudVolumegroupsGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVolumegroupsGetallOK struct {
	Payload *models.VolumeGroups
}

// IsSuccess returns true when this pcloud volumegroups getall o k response has a 2xx status code
func (o *PcloudVolumegroupsGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud volumegroups getall o k response has a 3xx status code
func (o *PcloudVolumegroupsGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups getall o k response has a 4xx status code
func (o *PcloudVolumegroupsGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volumegroups getall o k response has a 5xx status code
func (o *PcloudVolumegroupsGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups getall o k response a status code equal to that given
func (o *PcloudVolumegroupsGetallOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudVolumegroupsGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups][%d] pcloudVolumegroupsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudVolumegroupsGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups][%d] pcloudVolumegroupsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudVolumegroupsGetallOK) GetPayload() *models.VolumeGroups {
	return o.Payload
}

func (o *PcloudVolumegroupsGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroups)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetallBadRequest creates a PcloudVolumegroupsGetallBadRequest with default headers values
func NewPcloudVolumegroupsGetallBadRequest() *PcloudVolumegroupsGetallBadRequest {
	return &PcloudVolumegroupsGetallBadRequest{}
}

/*
PcloudVolumegroupsGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVolumegroupsGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups getall bad request response has a 2xx status code
func (o *PcloudVolumegroupsGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups getall bad request response has a 3xx status code
func (o *PcloudVolumegroupsGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups getall bad request response has a 4xx status code
func (o *PcloudVolumegroupsGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups getall bad request response has a 5xx status code
func (o *PcloudVolumegroupsGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups getall bad request response a status code equal to that given
func (o *PcloudVolumegroupsGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudVolumegroupsGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups][%d] pcloudVolumegroupsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVolumegroupsGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups][%d] pcloudVolumegroupsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVolumegroupsGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetallUnauthorized creates a PcloudVolumegroupsGetallUnauthorized with default headers values
func NewPcloudVolumegroupsGetallUnauthorized() *PcloudVolumegroupsGetallUnauthorized {
	return &PcloudVolumegroupsGetallUnauthorized{}
}

/*
PcloudVolumegroupsGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVolumegroupsGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups getall unauthorized response has a 2xx status code
func (o *PcloudVolumegroupsGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups getall unauthorized response has a 3xx status code
func (o *PcloudVolumegroupsGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups getall unauthorized response has a 4xx status code
func (o *PcloudVolumegroupsGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups getall unauthorized response has a 5xx status code
func (o *PcloudVolumegroupsGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups getall unauthorized response a status code equal to that given
func (o *PcloudVolumegroupsGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudVolumegroupsGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups][%d] pcloudVolumegroupsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVolumegroupsGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups][%d] pcloudVolumegroupsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVolumegroupsGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetallForbidden creates a PcloudVolumegroupsGetallForbidden with default headers values
func NewPcloudVolumegroupsGetallForbidden() *PcloudVolumegroupsGetallForbidden {
	return &PcloudVolumegroupsGetallForbidden{}
}

/*
PcloudVolumegroupsGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVolumegroupsGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups getall forbidden response has a 2xx status code
func (o *PcloudVolumegroupsGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups getall forbidden response has a 3xx status code
func (o *PcloudVolumegroupsGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups getall forbidden response has a 4xx status code
func (o *PcloudVolumegroupsGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups getall forbidden response has a 5xx status code
func (o *PcloudVolumegroupsGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups getall forbidden response a status code equal to that given
func (o *PcloudVolumegroupsGetallForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudVolumegroupsGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups][%d] pcloudVolumegroupsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVolumegroupsGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups][%d] pcloudVolumegroupsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVolumegroupsGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetallInternalServerError creates a PcloudVolumegroupsGetallInternalServerError with default headers values
func NewPcloudVolumegroupsGetallInternalServerError() *PcloudVolumegroupsGetallInternalServerError {
	return &PcloudVolumegroupsGetallInternalServerError{}
}

/*
PcloudVolumegroupsGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVolumegroupsGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups getall internal server error response has a 2xx status code
func (o *PcloudVolumegroupsGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups getall internal server error response has a 3xx status code
func (o *PcloudVolumegroupsGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups getall internal server error response has a 4xx status code
func (o *PcloudVolumegroupsGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volumegroups getall internal server error response has a 5xx status code
func (o *PcloudVolumegroupsGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud volumegroups getall internal server error response a status code equal to that given
func (o *PcloudVolumegroupsGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudVolumegroupsGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups][%d] pcloudVolumegroupsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVolumegroupsGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups][%d] pcloudVolumegroupsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVolumegroupsGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

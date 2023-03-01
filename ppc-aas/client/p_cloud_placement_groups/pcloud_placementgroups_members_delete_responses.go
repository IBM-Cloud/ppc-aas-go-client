// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_placement_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudPlacementgroupsMembersDeleteReader is a Reader for the PcloudPlacementgroupsMembersDelete structure.
type PcloudPlacementgroupsMembersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPlacementgroupsMembersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPlacementgroupsMembersDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPlacementgroupsMembersDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPlacementgroupsMembersDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPlacementgroupsMembersDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudPlacementgroupsMembersDeleteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudPlacementgroupsMembersDeleteUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPlacementgroupsMembersDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPlacementgroupsMembersDeleteOK creates a PcloudPlacementgroupsMembersDeleteOK with default headers values
func NewPcloudPlacementgroupsMembersDeleteOK() *PcloudPlacementgroupsMembersDeleteOK {
	return &PcloudPlacementgroupsMembersDeleteOK{}
}

/*
PcloudPlacementgroupsMembersDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPlacementgroupsMembersDeleteOK struct {
	Payload *models.PlacementGroup
}

// IsSuccess returns true when this pcloud placementgroups members delete o k response has a 2xx status code
func (o *PcloudPlacementgroupsMembersDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud placementgroups members delete o k response has a 3xx status code
func (o *PcloudPlacementgroupsMembersDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members delete o k response has a 4xx status code
func (o *PcloudPlacementgroupsMembersDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud placementgroups members delete o k response has a 5xx status code
func (o *PcloudPlacementgroupsMembersDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups members delete o k response a status code equal to that given
func (o *PcloudPlacementgroupsMembersDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudPlacementgroupsMembersDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteOK) GetPayload() *models.PlacementGroup {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlacementGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersDeleteBadRequest creates a PcloudPlacementgroupsMembersDeleteBadRequest with default headers values
func NewPcloudPlacementgroupsMembersDeleteBadRequest() *PcloudPlacementgroupsMembersDeleteBadRequest {
	return &PcloudPlacementgroupsMembersDeleteBadRequest{}
}

/*
PcloudPlacementgroupsMembersDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPlacementgroupsMembersDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups members delete bad request response has a 2xx status code
func (o *PcloudPlacementgroupsMembersDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups members delete bad request response has a 3xx status code
func (o *PcloudPlacementgroupsMembersDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members delete bad request response has a 4xx status code
func (o *PcloudPlacementgroupsMembersDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups members delete bad request response has a 5xx status code
func (o *PcloudPlacementgroupsMembersDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups members delete bad request response a status code equal to that given
func (o *PcloudPlacementgroupsMembersDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudPlacementgroupsMembersDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersDeleteForbidden creates a PcloudPlacementgroupsMembersDeleteForbidden with default headers values
func NewPcloudPlacementgroupsMembersDeleteForbidden() *PcloudPlacementgroupsMembersDeleteForbidden {
	return &PcloudPlacementgroupsMembersDeleteForbidden{}
}

/*
PcloudPlacementgroupsMembersDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPlacementgroupsMembersDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups members delete forbidden response has a 2xx status code
func (o *PcloudPlacementgroupsMembersDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups members delete forbidden response has a 3xx status code
func (o *PcloudPlacementgroupsMembersDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members delete forbidden response has a 4xx status code
func (o *PcloudPlacementgroupsMembersDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups members delete forbidden response has a 5xx status code
func (o *PcloudPlacementgroupsMembersDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups members delete forbidden response a status code equal to that given
func (o *PcloudPlacementgroupsMembersDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudPlacementgroupsMembersDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersDeleteNotFound creates a PcloudPlacementgroupsMembersDeleteNotFound with default headers values
func NewPcloudPlacementgroupsMembersDeleteNotFound() *PcloudPlacementgroupsMembersDeleteNotFound {
	return &PcloudPlacementgroupsMembersDeleteNotFound{}
}

/*
PcloudPlacementgroupsMembersDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPlacementgroupsMembersDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups members delete not found response has a 2xx status code
func (o *PcloudPlacementgroupsMembersDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups members delete not found response has a 3xx status code
func (o *PcloudPlacementgroupsMembersDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members delete not found response has a 4xx status code
func (o *PcloudPlacementgroupsMembersDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups members delete not found response has a 5xx status code
func (o *PcloudPlacementgroupsMembersDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups members delete not found response a status code equal to that given
func (o *PcloudPlacementgroupsMembersDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudPlacementgroupsMembersDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersDeleteConflict creates a PcloudPlacementgroupsMembersDeleteConflict with default headers values
func NewPcloudPlacementgroupsMembersDeleteConflict() *PcloudPlacementgroupsMembersDeleteConflict {
	return &PcloudPlacementgroupsMembersDeleteConflict{}
}

/*
PcloudPlacementgroupsMembersDeleteConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudPlacementgroupsMembersDeleteConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups members delete conflict response has a 2xx status code
func (o *PcloudPlacementgroupsMembersDeleteConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups members delete conflict response has a 3xx status code
func (o *PcloudPlacementgroupsMembersDeleteConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members delete conflict response has a 4xx status code
func (o *PcloudPlacementgroupsMembersDeleteConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups members delete conflict response has a 5xx status code
func (o *PcloudPlacementgroupsMembersDeleteConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups members delete conflict response a status code equal to that given
func (o *PcloudPlacementgroupsMembersDeleteConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PcloudPlacementgroupsMembersDeleteConflict) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteConflict  %+v", 409, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteConflict) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteConflict  %+v", 409, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersDeleteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersDeleteUnprocessableEntity creates a PcloudPlacementgroupsMembersDeleteUnprocessableEntity with default headers values
func NewPcloudPlacementgroupsMembersDeleteUnprocessableEntity() *PcloudPlacementgroupsMembersDeleteUnprocessableEntity {
	return &PcloudPlacementgroupsMembersDeleteUnprocessableEntity{}
}

/*
PcloudPlacementgroupsMembersDeleteUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudPlacementgroupsMembersDeleteUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups members delete unprocessable entity response has a 2xx status code
func (o *PcloudPlacementgroupsMembersDeleteUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups members delete unprocessable entity response has a 3xx status code
func (o *PcloudPlacementgroupsMembersDeleteUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members delete unprocessable entity response has a 4xx status code
func (o *PcloudPlacementgroupsMembersDeleteUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups members delete unprocessable entity response has a 5xx status code
func (o *PcloudPlacementgroupsMembersDeleteUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups members delete unprocessable entity response a status code equal to that given
func (o *PcloudPlacementgroupsMembersDeleteUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PcloudPlacementgroupsMembersDeleteUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersDeleteUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersDeleteInternalServerError creates a PcloudPlacementgroupsMembersDeleteInternalServerError with default headers values
func NewPcloudPlacementgroupsMembersDeleteInternalServerError() *PcloudPlacementgroupsMembersDeleteInternalServerError {
	return &PcloudPlacementgroupsMembersDeleteInternalServerError{}
}

/*
PcloudPlacementgroupsMembersDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPlacementgroupsMembersDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups members delete internal server error response has a 2xx status code
func (o *PcloudPlacementgroupsMembersDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups members delete internal server error response has a 3xx status code
func (o *PcloudPlacementgroupsMembersDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members delete internal server error response has a 4xx status code
func (o *PcloudPlacementgroupsMembersDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud placementgroups members delete internal server error response has a 5xx status code
func (o *PcloudPlacementgroupsMembersDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud placementgroups members delete internal server error response a status code equal to that given
func (o *PcloudPlacementgroupsMembersDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudPlacementgroupsMembersDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPlacementgroupsMembersDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_s_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudSapGetallReader is a Reader for the PcloudSapGetall structure.
type PcloudSapGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudSapGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudSapGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudSapGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudSapGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudSapGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap] pcloud.sap.getall", response, response.Code())
	}
}

// NewPcloudSapGetallOK creates a PcloudSapGetallOK with default headers values
func NewPcloudSapGetallOK() *PcloudSapGetallOK {
	return &PcloudSapGetallOK{}
}

/*
PcloudSapGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudSapGetallOK struct {
	Payload *models.SAPProfiles
}

// IsSuccess returns true when this pcloud sap getall o k response has a 2xx status code
func (o *PcloudSapGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud sap getall o k response has a 3xx status code
func (o *PcloudSapGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap getall o k response has a 4xx status code
func (o *PcloudSapGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sap getall o k response has a 5xx status code
func (o *PcloudSapGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap getall o k response a status code equal to that given
func (o *PcloudSapGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud sap getall o k response
func (o *PcloudSapGetallOK) Code() int {
	return 200
}

func (o *PcloudSapGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudSapGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudSapGetallOK) GetPayload() *models.SAPProfiles {
	return o.Payload
}

func (o *PcloudSapGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SAPProfiles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapGetallBadRequest creates a PcloudSapGetallBadRequest with default headers values
func NewPcloudSapGetallBadRequest() *PcloudSapGetallBadRequest {
	return &PcloudSapGetallBadRequest{}
}

/*
PcloudSapGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudSapGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sap getall bad request response has a 2xx status code
func (o *PcloudSapGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sap getall bad request response has a 3xx status code
func (o *PcloudSapGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap getall bad request response has a 4xx status code
func (o *PcloudSapGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sap getall bad request response has a 5xx status code
func (o *PcloudSapGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap getall bad request response a status code equal to that given
func (o *PcloudSapGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud sap getall bad request response
func (o *PcloudSapGetallBadRequest) Code() int {
	return 400
}

func (o *PcloudSapGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudSapGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudSapGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSapGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapGetallUnauthorized creates a PcloudSapGetallUnauthorized with default headers values
func NewPcloudSapGetallUnauthorized() *PcloudSapGetallUnauthorized {
	return &PcloudSapGetallUnauthorized{}
}

/*
PcloudSapGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudSapGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sap getall unauthorized response has a 2xx status code
func (o *PcloudSapGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sap getall unauthorized response has a 3xx status code
func (o *PcloudSapGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap getall unauthorized response has a 4xx status code
func (o *PcloudSapGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sap getall unauthorized response has a 5xx status code
func (o *PcloudSapGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap getall unauthorized response a status code equal to that given
func (o *PcloudSapGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud sap getall unauthorized response
func (o *PcloudSapGetallUnauthorized) Code() int {
	return 401
}

func (o *PcloudSapGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudSapGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudSapGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSapGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapGetallInternalServerError creates a PcloudSapGetallInternalServerError with default headers values
func NewPcloudSapGetallInternalServerError() *PcloudSapGetallInternalServerError {
	return &PcloudSapGetallInternalServerError{}
}

/*
PcloudSapGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudSapGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sap getall internal server error response has a 2xx status code
func (o *PcloudSapGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sap getall internal server error response has a 3xx status code
func (o *PcloudSapGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap getall internal server error response has a 4xx status code
func (o *PcloudSapGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sap getall internal server error response has a 5xx status code
func (o *PcloudSapGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud sap getall internal server error response a status code equal to that given
func (o *PcloudSapGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud sap getall internal server error response
func (o *PcloudSapGetallInternalServerError) Code() int {
	return 500
}

func (o *PcloudSapGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudSapGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudSapGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSapGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

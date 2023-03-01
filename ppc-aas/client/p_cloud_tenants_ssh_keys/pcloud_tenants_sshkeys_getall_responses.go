// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tenants_ssh_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudTenantsSshkeysGetallReader is a Reader for the PcloudTenantsSshkeysGetall structure.
type PcloudTenantsSshkeysGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudTenantsSshkeysGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudTenantsSshkeysGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudTenantsSshkeysGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudTenantsSshkeysGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudTenantsSshkeysGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudTenantsSshkeysGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudTenantsSshkeysGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudTenantsSshkeysGetallOK creates a PcloudTenantsSshkeysGetallOK with default headers values
func NewPcloudTenantsSshkeysGetallOK() *PcloudTenantsSshkeysGetallOK {
	return &PcloudTenantsSshkeysGetallOK{}
}

/*
PcloudTenantsSshkeysGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudTenantsSshkeysGetallOK struct {
	Payload *models.SSHKeys
}

// IsSuccess returns true when this pcloud tenants sshkeys getall o k response has a 2xx status code
func (o *PcloudTenantsSshkeysGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud tenants sshkeys getall o k response has a 3xx status code
func (o *PcloudTenantsSshkeysGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys getall o k response has a 4xx status code
func (o *PcloudTenantsSshkeysGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud tenants sshkeys getall o k response has a 5xx status code
func (o *PcloudTenantsSshkeysGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants sshkeys getall o k response a status code equal to that given
func (o *PcloudTenantsSshkeysGetallOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudTenantsSshkeysGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudTenantsSshkeysGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudTenantsSshkeysGetallOK) GetPayload() *models.SSHKeys {
	return o.Payload
}

func (o *PcloudTenantsSshkeysGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHKeys)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysGetallBadRequest creates a PcloudTenantsSshkeysGetallBadRequest with default headers values
func NewPcloudTenantsSshkeysGetallBadRequest() *PcloudTenantsSshkeysGetallBadRequest {
	return &PcloudTenantsSshkeysGetallBadRequest{}
}

/*
PcloudTenantsSshkeysGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudTenantsSshkeysGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants sshkeys getall bad request response has a 2xx status code
func (o *PcloudTenantsSshkeysGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants sshkeys getall bad request response has a 3xx status code
func (o *PcloudTenantsSshkeysGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys getall bad request response has a 4xx status code
func (o *PcloudTenantsSshkeysGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants sshkeys getall bad request response has a 5xx status code
func (o *PcloudTenantsSshkeysGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants sshkeys getall bad request response a status code equal to that given
func (o *PcloudTenantsSshkeysGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudTenantsSshkeysGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTenantsSshkeysGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTenantsSshkeysGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysGetallUnauthorized creates a PcloudTenantsSshkeysGetallUnauthorized with default headers values
func NewPcloudTenantsSshkeysGetallUnauthorized() *PcloudTenantsSshkeysGetallUnauthorized {
	return &PcloudTenantsSshkeysGetallUnauthorized{}
}

/*
PcloudTenantsSshkeysGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudTenantsSshkeysGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants sshkeys getall unauthorized response has a 2xx status code
func (o *PcloudTenantsSshkeysGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants sshkeys getall unauthorized response has a 3xx status code
func (o *PcloudTenantsSshkeysGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys getall unauthorized response has a 4xx status code
func (o *PcloudTenantsSshkeysGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants sshkeys getall unauthorized response has a 5xx status code
func (o *PcloudTenantsSshkeysGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants sshkeys getall unauthorized response a status code equal to that given
func (o *PcloudTenantsSshkeysGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudTenantsSshkeysGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudTenantsSshkeysGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudTenantsSshkeysGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysGetallForbidden creates a PcloudTenantsSshkeysGetallForbidden with default headers values
func NewPcloudTenantsSshkeysGetallForbidden() *PcloudTenantsSshkeysGetallForbidden {
	return &PcloudTenantsSshkeysGetallForbidden{}
}

/*
PcloudTenantsSshkeysGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudTenantsSshkeysGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants sshkeys getall forbidden response has a 2xx status code
func (o *PcloudTenantsSshkeysGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants sshkeys getall forbidden response has a 3xx status code
func (o *PcloudTenantsSshkeysGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys getall forbidden response has a 4xx status code
func (o *PcloudTenantsSshkeysGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants sshkeys getall forbidden response has a 5xx status code
func (o *PcloudTenantsSshkeysGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants sshkeys getall forbidden response a status code equal to that given
func (o *PcloudTenantsSshkeysGetallForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudTenantsSshkeysGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudTenantsSshkeysGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudTenantsSshkeysGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysGetallNotFound creates a PcloudTenantsSshkeysGetallNotFound with default headers values
func NewPcloudTenantsSshkeysGetallNotFound() *PcloudTenantsSshkeysGetallNotFound {
	return &PcloudTenantsSshkeysGetallNotFound{}
}

/*
PcloudTenantsSshkeysGetallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudTenantsSshkeysGetallNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants sshkeys getall not found response has a 2xx status code
func (o *PcloudTenantsSshkeysGetallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants sshkeys getall not found response has a 3xx status code
func (o *PcloudTenantsSshkeysGetallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys getall not found response has a 4xx status code
func (o *PcloudTenantsSshkeysGetallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants sshkeys getall not found response has a 5xx status code
func (o *PcloudTenantsSshkeysGetallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants sshkeys getall not found response a status code equal to that given
func (o *PcloudTenantsSshkeysGetallNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudTenantsSshkeysGetallNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudTenantsSshkeysGetallNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudTenantsSshkeysGetallNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysGetallInternalServerError creates a PcloudTenantsSshkeysGetallInternalServerError with default headers values
func NewPcloudTenantsSshkeysGetallInternalServerError() *PcloudTenantsSshkeysGetallInternalServerError {
	return &PcloudTenantsSshkeysGetallInternalServerError{}
}

/*
PcloudTenantsSshkeysGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudTenantsSshkeysGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants sshkeys getall internal server error response has a 2xx status code
func (o *PcloudTenantsSshkeysGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants sshkeys getall internal server error response has a 3xx status code
func (o *PcloudTenantsSshkeysGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys getall internal server error response has a 4xx status code
func (o *PcloudTenantsSshkeysGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud tenants sshkeys getall internal server error response has a 5xx status code
func (o *PcloudTenantsSshkeysGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud tenants sshkeys getall internal server error response a status code equal to that given
func (o *PcloudTenantsSshkeysGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudTenantsSshkeysGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTenantsSshkeysGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}/sshkeys][%d] pcloudTenantsSshkeysGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTenantsSshkeysGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

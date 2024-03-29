// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_service_d_h_c_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudDhcpGetallReader is a Reader for the PcloudDhcpGetall structure.
type PcloudDhcpGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudDhcpGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudDhcpGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPcloudDhcpGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudDhcpGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp] pcloud.dhcp.getall", response, response.Code())
	}
}

// NewPcloudDhcpGetallOK creates a PcloudDhcpGetallOK with default headers values
func NewPcloudDhcpGetallOK() *PcloudDhcpGetallOK {
	return &PcloudDhcpGetallOK{}
}

/*
PcloudDhcpGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudDhcpGetallOK struct {
	Payload models.DHCPServers
}

// IsSuccess returns true when this pcloud dhcp getall o k response has a 2xx status code
func (o *PcloudDhcpGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud dhcp getall o k response has a 3xx status code
func (o *PcloudDhcpGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp getall o k response has a 4xx status code
func (o *PcloudDhcpGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud dhcp getall o k response has a 5xx status code
func (o *PcloudDhcpGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp getall o k response a status code equal to that given
func (o *PcloudDhcpGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud dhcp getall o k response
func (o *PcloudDhcpGetallOK) Code() int {
	return 200
}

func (o *PcloudDhcpGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudDhcpGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudDhcpGetallOK) GetPayload() models.DHCPServers {
	return o.Payload
}

func (o *PcloudDhcpGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpGetallForbidden creates a PcloudDhcpGetallForbidden with default headers values
func NewPcloudDhcpGetallForbidden() *PcloudDhcpGetallForbidden {
	return &PcloudDhcpGetallForbidden{}
}

/*
PcloudDhcpGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudDhcpGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp getall forbidden response has a 2xx status code
func (o *PcloudDhcpGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp getall forbidden response has a 3xx status code
func (o *PcloudDhcpGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp getall forbidden response has a 4xx status code
func (o *PcloudDhcpGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud dhcp getall forbidden response has a 5xx status code
func (o *PcloudDhcpGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp getall forbidden response a status code equal to that given
func (o *PcloudDhcpGetallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud dhcp getall forbidden response
func (o *PcloudDhcpGetallForbidden) Code() int {
	return 403
}

func (o *PcloudDhcpGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudDhcpGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudDhcpGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpGetallInternalServerError creates a PcloudDhcpGetallInternalServerError with default headers values
func NewPcloudDhcpGetallInternalServerError() *PcloudDhcpGetallInternalServerError {
	return &PcloudDhcpGetallInternalServerError{}
}

/*
PcloudDhcpGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudDhcpGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp getall internal server error response has a 2xx status code
func (o *PcloudDhcpGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp getall internal server error response has a 3xx status code
func (o *PcloudDhcpGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp getall internal server error response has a 4xx status code
func (o *PcloudDhcpGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud dhcp getall internal server error response has a 5xx status code
func (o *PcloudDhcpGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud dhcp getall internal server error response a status code equal to that given
func (o *PcloudDhcpGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud dhcp getall internal server error response
func (o *PcloudDhcpGetallInternalServerError) Code() int {
	return 500
}

func (o *PcloudDhcpGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudDhcpGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudDhcpGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

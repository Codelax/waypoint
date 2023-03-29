// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/waypoint/pkg/client/gen/models"
)

// WaypointListInstances2Reader is a Reader for the WaypointListInstances2 structure.
type WaypointListInstances2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListInstances2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListInstances2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListInstances2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListInstances2OK creates a WaypointListInstances2OK with default headers values
func NewWaypointListInstances2OK() *WaypointListInstances2OK {
	return &WaypointListInstances2OK{}
}

/*
WaypointListInstances2OK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointListInstances2OK struct {
	Payload *models.HashicorpWaypointListInstancesResponse
}

// IsSuccess returns true when this waypoint list instances2 o k response has a 2xx status code
func (o *WaypointListInstances2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint list instances2 o k response has a 3xx status code
func (o *WaypointListInstances2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint list instances2 o k response has a 4xx status code
func (o *WaypointListInstances2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint list instances2 o k response has a 5xx status code
func (o *WaypointListInstances2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint list instances2 o k response a status code equal to that given
func (o *WaypointListInstances2OK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointListInstances2OK) Error() string {
	return fmt.Sprintf("[GET /project/{application.application.project}/application/{application.application.application}/instances][%d] waypointListInstances2OK  %+v", 200, o.Payload)
}

func (o *WaypointListInstances2OK) String() string {
	return fmt.Sprintf("[GET /project/{application.application.project}/application/{application.application.application}/instances][%d] waypointListInstances2OK  %+v", 200, o.Payload)
}

func (o *WaypointListInstances2OK) GetPayload() *models.HashicorpWaypointListInstancesResponse {
	return o.Payload
}

func (o *WaypointListInstances2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListInstancesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListInstances2Default creates a WaypointListInstances2Default with default headers values
func NewWaypointListInstances2Default(code int) *WaypointListInstances2Default {
	return &WaypointListInstances2Default{
		_statusCode: code,
	}
}

/*
WaypointListInstances2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointListInstances2Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list instances2 default response
func (o *WaypointListInstances2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint list instances2 default response has a 2xx status code
func (o *WaypointListInstances2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint list instances2 default response has a 3xx status code
func (o *WaypointListInstances2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint list instances2 default response has a 4xx status code
func (o *WaypointListInstances2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint list instances2 default response has a 5xx status code
func (o *WaypointListInstances2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint list instances2 default response a status code equal to that given
func (o *WaypointListInstances2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointListInstances2Default) Error() string {
	return fmt.Sprintf("[GET /project/{application.application.project}/application/{application.application.application}/instances][%d] Waypoint_ListInstances2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListInstances2Default) String() string {
	return fmt.Sprintf("[GET /project/{application.application.project}/application/{application.application.application}/instances][%d] Waypoint_ListInstances2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListInstances2Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListInstances2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
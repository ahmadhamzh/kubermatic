// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// DeleteSeedReader is a Reader for the DeleteSeed structure.
type DeleteSeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSeedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteSeedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSeedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSeedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSeedOK creates a DeleteSeedOK with default headers values
func NewDeleteSeedOK() *DeleteSeedOK {
	return &DeleteSeedOK{}
}

/*DeleteSeedOK handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteSeedOK struct {
}

func (o *DeleteSeedOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}][%d] deleteSeedOK ", 200)
}

func (o *DeleteSeedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSeedUnauthorized creates a DeleteSeedUnauthorized with default headers values
func NewDeleteSeedUnauthorized() *DeleteSeedUnauthorized {
	return &DeleteSeedUnauthorized{}
}

/*DeleteSeedUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteSeedUnauthorized struct {
}

func (o *DeleteSeedUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}][%d] deleteSeedUnauthorized ", 401)
}

func (o *DeleteSeedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSeedForbidden creates a DeleteSeedForbidden with default headers values
func NewDeleteSeedForbidden() *DeleteSeedForbidden {
	return &DeleteSeedForbidden{}
}

/*DeleteSeedForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteSeedForbidden struct {
}

func (o *DeleteSeedForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}][%d] deleteSeedForbidden ", 403)
}

func (o *DeleteSeedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSeedDefault creates a DeleteSeedDefault with default headers values
func NewDeleteSeedDefault(code int) *DeleteSeedDefault {
	return &DeleteSeedDefault{
		_statusCode: code,
	}
}

/*DeleteSeedDefault handles this case with default header values.

errorResponse
*/
type DeleteSeedDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete seed default response
func (o *DeleteSeedDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSeedDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/seeds/{seed_name}][%d] deleteSeed default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSeedDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteSeedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rackn/rocket-skates/models"
)

// PostFileReader is a Reader for the PostFile structure.
type PostFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostFileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostFileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 507:
		result := NewPostFileInsufficientStorage()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostFileCreated creates a PostFileCreated with default headers values
func NewPostFileCreated() *PostFileCreated {
	return &PostFileCreated{}
}

/*PostFileCreated handles this case with default header values.

PostFileCreated post file created
*/
type PostFileCreated struct {
	Payload PostFileCreatedBody
}

func (o *PostFileCreated) Error() string {
	return fmt.Sprintf("[POST /files/{path}][%d] postFileCreated  %+v", 201, o.Payload)
}

func (o *PostFileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFileBadRequest creates a PostFileBadRequest with default headers values
func NewPostFileBadRequest() *PostFileBadRequest {
	return &PostFileBadRequest{}
}

/*PostFileBadRequest handles this case with default header values.

PostFileBadRequest post file bad request
*/
type PostFileBadRequest struct {
	Payload *models.Result
}

func (o *PostFileBadRequest) Error() string {
	return fmt.Sprintf("[POST /files/{path}][%d] postFileBadRequest  %+v", 400, o.Payload)
}

func (o *PostFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFileConflict creates a PostFileConflict with default headers values
func NewPostFileConflict() *PostFileConflict {
	return &PostFileConflict{}
}

/*PostFileConflict handles this case with default header values.

PostFileConflict post file conflict
*/
type PostFileConflict struct {
	Payload *models.Result
}

func (o *PostFileConflict) Error() string {
	return fmt.Sprintf("[POST /files/{path}][%d] postFileConflict  %+v", 409, o.Payload)
}

func (o *PostFileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFileInsufficientStorage creates a PostFileInsufficientStorage with default headers values
func NewPostFileInsufficientStorage() *PostFileInsufficientStorage {
	return &PostFileInsufficientStorage{}
}

/*PostFileInsufficientStorage handles this case with default header values.

PostFileInsufficientStorage post file insufficient storage
*/
type PostFileInsufficientStorage struct {
	Payload *models.Result
}

func (o *PostFileInsufficientStorage) Error() string {
	return fmt.Sprintf("[POST /files/{path}][%d] postFileInsufficientStorage  %+v", 507, o.Payload)
}

func (o *PostFileInsufficientStorage) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostFileCreatedBody post file created body
swagger:model PostFileCreatedBody
*/
type PostFileCreatedBody struct {

	// name
	// Required: true
	Name *string `json:"Name"`

	// size
	// Required: true
	Size *int64 `json:"Size"`
}
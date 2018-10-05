// Code generated by go-swagger; DO NOT EDIT.

package scrape_configs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/percona/pmm-managed/api/swagger/models"
)

// ListMixin4Reader is a Reader for the ListMixin4 structure.
type ListMixin4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMixin4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListMixin4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListMixin4OK creates a ListMixin4OK with default headers values
func NewListMixin4OK() *ListMixin4OK {
	return &ListMixin4OK{}
}

/*ListMixin4OK handles this case with default header values.

(empty)
*/
type ListMixin4OK struct {
	Payload *models.APIScrapeConfigsListResponse
}

func (o *ListMixin4OK) Error() string {
	return fmt.Sprintf("[GET /v0/scrape-configs][%d] listMixin4OK  %+v", 200, o.Payload)
}

func (o *ListMixin4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIScrapeConfigsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

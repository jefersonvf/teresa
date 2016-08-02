package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/teresa/api/models"
)

// NewCreateAppParams creates a new CreateAppParams object
// with the default values initialized.
func NewCreateAppParams() *CreateAppParams {
	var ()
	return &CreateAppParams{}
}

/*CreateAppParams contains all the parameters to send to the API endpoint
for the create app operation typically these are written to a http.Request
*/
type CreateAppParams struct {

	/*Body*/
	Body *models.App
	/*TeamID
	  Team ID

	*/
	TeamID int64
}

// WithBody adds the body to the create app params
func (o *CreateAppParams) WithBody(Body *models.App) *CreateAppParams {
	o.Body = Body
	return o
}

// WithTeamID adds the teamId to the create app params
func (o *CreateAppParams) WithTeamID(TeamID int64) *CreateAppParams {
	o.TeamID = TeamID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.App)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param team_id
	if err := r.SetPathParam("team_id", swag.FormatInt64(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

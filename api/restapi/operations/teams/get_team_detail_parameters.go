package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTeamDetailParams creates a new GetTeamDetailParams object
// with the default values initialized.
func NewGetTeamDetailParams() GetTeamDetailParams {
	var ()
	return GetTeamDetailParams{}
}

// GetTeamDetailParams contains all the bound params for the get team detail operation
// typically these are obtained from a http.Request
//
// swagger:parameters getTeamDetail
type GetTeamDetailParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Team ID
	  Required: true
	  In: path
	*/
	TeamID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetTeamDetailParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rTeamID, rhkTeamID, _ := route.Params.GetOK("team_id")
	if err := o.bindTeamID(rTeamID, rhkTeamID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTeamDetailParams) bindTeamID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("team_id", "path", "int64", raw)
	}
	o.TeamID = value

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListNodesForClusterLegacyParams creates a new ListNodesForClusterLegacyParams object
// with the default values initialized.
func NewListNodesForClusterLegacyParams() *ListNodesForClusterLegacyParams {
	var ()
	return &ListNodesForClusterLegacyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListNodesForClusterLegacyParamsWithTimeout creates a new ListNodesForClusterLegacyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListNodesForClusterLegacyParamsWithTimeout(timeout time.Duration) *ListNodesForClusterLegacyParams {
	var ()
	return &ListNodesForClusterLegacyParams{

		timeout: timeout,
	}
}

// NewListNodesForClusterLegacyParamsWithContext creates a new ListNodesForClusterLegacyParams object
// with the default values initialized, and the ability to set a context for a request
func NewListNodesForClusterLegacyParamsWithContext(ctx context.Context) *ListNodesForClusterLegacyParams {
	var ()
	return &ListNodesForClusterLegacyParams{

		Context: ctx,
	}
}

// NewListNodesForClusterLegacyParamsWithHTTPClient creates a new ListNodesForClusterLegacyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListNodesForClusterLegacyParamsWithHTTPClient(client *http.Client) *ListNodesForClusterLegacyParams {
	var ()
	return &ListNodesForClusterLegacyParams{
		HTTPClient: client,
	}
}

/*ListNodesForClusterLegacyParams contains all the parameters to send to the API endpoint
for the list nodes for cluster legacy operation typically these are written to a http.Request
*/
type ListNodesForClusterLegacyParams struct {

	/*ClusterID*/
	ClusterID string
	/*Dc*/
	Dc string
	/*HideInitialConditions*/
	HideInitialConditions *bool
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) WithTimeout(timeout time.Duration) *ListNodesForClusterLegacyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) WithContext(ctx context.Context) *ListNodesForClusterLegacyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) WithHTTPClient(client *http.Client) *ListNodesForClusterLegacyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) WithClusterID(clusterID string) *ListNodesForClusterLegacyParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDc adds the dc to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) WithDc(dc string) *ListNodesForClusterLegacyParams {
	o.SetDc(dc)
	return o
}

// SetDc adds the dc to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) SetDc(dc string) {
	o.Dc = dc
}

// WithHideInitialConditions adds the hideInitialConditions to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) WithHideInitialConditions(hideInitialConditions *bool) *ListNodesForClusterLegacyParams {
	o.SetHideInitialConditions(hideInitialConditions)
	return o
}

// SetHideInitialConditions adds the hideInitialConditions to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) SetHideInitialConditions(hideInitialConditions *bool) {
	o.HideInitialConditions = hideInitialConditions
}

// WithProjectID adds the projectID to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) WithProjectID(projectID string) *ListNodesForClusterLegacyParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list nodes for cluster legacy params
func (o *ListNodesForClusterLegacyParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListNodesForClusterLegacyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.Dc); err != nil {
		return err
	}

	if o.HideInitialConditions != nil {

		// query param hideInitialConditions
		var qrHideInitialConditions bool
		if o.HideInitialConditions != nil {
			qrHideInitialConditions = *o.HideInitialConditions
		}
		qHideInitialConditions := swag.FormatBool(qrHideInitialConditions)
		if qHideInitialConditions != "" {
			if err := r.SetQueryParam("hideInitialConditions", qHideInitialConditions); err != nil {
				return err
			}
		}

	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

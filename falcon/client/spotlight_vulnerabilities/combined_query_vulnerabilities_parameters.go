// Code generated by go-swagger; DO NOT EDIT.

package spotlight_vulnerabilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewCombinedQueryVulnerabilitiesParams creates a new CombinedQueryVulnerabilitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCombinedQueryVulnerabilitiesParams() *CombinedQueryVulnerabilitiesParams {
	return &CombinedQueryVulnerabilitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCombinedQueryVulnerabilitiesParamsWithTimeout creates a new CombinedQueryVulnerabilitiesParams object
// with the ability to set a timeout on a request.
func NewCombinedQueryVulnerabilitiesParamsWithTimeout(timeout time.Duration) *CombinedQueryVulnerabilitiesParams {
	return &CombinedQueryVulnerabilitiesParams{
		timeout: timeout,
	}
}

// NewCombinedQueryVulnerabilitiesParamsWithContext creates a new CombinedQueryVulnerabilitiesParams object
// with the ability to set a context for a request.
func NewCombinedQueryVulnerabilitiesParamsWithContext(ctx context.Context) *CombinedQueryVulnerabilitiesParams {
	return &CombinedQueryVulnerabilitiesParams{
		Context: ctx,
	}
}

// NewCombinedQueryVulnerabilitiesParamsWithHTTPClient creates a new CombinedQueryVulnerabilitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCombinedQueryVulnerabilitiesParamsWithHTTPClient(client *http.Client) *CombinedQueryVulnerabilitiesParams {
	return &CombinedQueryVulnerabilitiesParams{
		HTTPClient: client,
	}
}

/* CombinedQueryVulnerabilitiesParams contains all the parameters to send to the API endpoint
   for the combined query vulnerabilities operation.

   Typically these are written to a http.Request.
*/
type CombinedQueryVulnerabilitiesParams struct {

	/* After.

	   A pagination token used with the `limit` parameter to manage pagination of results. On your first request, don't provide an `after` token. On subsequent requests, provide the `after` token from the previous response to continue from that place in the results.
	*/
	After *string

	/* Facet.

	     Select various details blocks to be returned for each vulnerability entity. Supported values:

	<ul><li>host_info</li><li>remediation_details</li><li>cve_details</li></ul>
	*/
	Facet []string

	/* Filter.

	     Filter items using a query in Falcon Query Language (FQL). Wildcards * are unsupported.

	Common filter options include:

	<ul><li>created_timestamp:>'2019-11-25T22:36:12Z'</li><li>closed_timestamp:>'2019-11-25T22:36:12Z'</li><li>aid:'8e7656b27d8c49a34a1af416424d6231'</li></ul>
	*/
	Filter string

	/* Limit.

	   The number of items to return in this response (default: 100, max: 5000). Use with the after parameter to manage pagination of results.
	*/
	Limit *int64

	/* Sort.

	     Sort vulnerabilities by their properties. Common sort options include:

	<ul><li>created_timestamp|desc</li><li>closed_timestamp|asc</li></ul>
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the combined query vulnerabilities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CombinedQueryVulnerabilitiesParams) WithDefaults() *CombinedQueryVulnerabilitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the combined query vulnerabilities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CombinedQueryVulnerabilitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) WithTimeout(timeout time.Duration) *CombinedQueryVulnerabilitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) WithContext(ctx context.Context) *CombinedQueryVulnerabilitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) WithHTTPClient(client *http.Client) *CombinedQueryVulnerabilitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) WithAfter(after *string) *CombinedQueryVulnerabilitiesParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) SetAfter(after *string) {
	o.After = after
}

// WithFacet adds the facet to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) WithFacet(facet []string) *CombinedQueryVulnerabilitiesParams {
	o.SetFacet(facet)
	return o
}

// SetFacet adds the facet to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) SetFacet(facet []string) {
	o.Facet = facet
}

// WithFilter adds the filter to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) WithFilter(filter string) *CombinedQueryVulnerabilitiesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) SetFilter(filter string) {
	o.Filter = filter
}

// WithLimit adds the limit to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) WithLimit(limit *int64) *CombinedQueryVulnerabilitiesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithSort adds the sort to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) WithSort(sort *string) *CombinedQueryVulnerabilitiesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the combined query vulnerabilities params
func (o *CombinedQueryVulnerabilitiesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *CombinedQueryVulnerabilitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter string

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	if o.Facet != nil {

		// binding items for facet
		joinedFacet := o.bindParamFacet(reg)

		// query array param facet
		if err := r.SetQueryParam("facet", joinedFacet...); err != nil {
			return err
		}
	}

	// query param filter
	qrFilter := o.Filter
	qFilter := qrFilter
	if qFilter != "" {

		if err := r.SetQueryParam("filter", qFilter); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCombinedQueryVulnerabilities binds the parameter facet
func (o *CombinedQueryVulnerabilitiesParams) bindParamFacet(formats strfmt.Registry) []string {
	facetIR := o.Facet

	var facetIC []string
	for _, facetIIR := range facetIR { // explode []string

		facetIIV := facetIIR // string as string
		facetIC = append(facetIC, facetIIV)
	}

	// items.CollectionFormat: "csv"
	facetIS := swag.JoinByFormat(facetIC, "csv")

	return facetIS
}

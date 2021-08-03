package searchservice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
)

// IndexersClient is the client that can be used to manage and query indexes and documents, as well as manage other
// resources, on a search service.
type IndexersClient struct {
	BaseClient
}

// NewIndexersClient creates an instance of the IndexersClient client.
func NewIndexersClient(searchServiceName string) IndexersClient {
	return IndexersClient{New(searchServiceName)}
}

// Create creates a new indexer.
// Parameters:
// indexer - the definition of the indexer to create.
// clientRequestID - the tracking ID sent with the request to help with debugging.
func (client IndexersClient) Create(ctx context.Context, indexer Indexer, clientRequestID *uuid.UUID) (result Indexer, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IndexersClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: indexer,
			Constraints: []validation.Constraint{{Target: "indexer.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "indexer.DataSourceName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "indexer.TargetIndexName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "indexer.Schedule", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "indexer.Schedule.Interval", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("searchservice.IndexersClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, indexer, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client IndexersClient) CreatePreparer(ctx context.Context, indexer Indexer, clientRequestID *uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPath("/indexers"),
		autorest.WithJSON(indexer),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client IndexersClient) CreateResponder(resp *http.Response) (result Indexer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates a new indexer or updates an indexer if it already exists.
// Parameters:
// indexerName - the name of the indexer to create or update.
// indexer - the definition of the indexer to create or update.
// clientRequestID - the tracking ID sent with the request to help with debugging.
// ifMatch - defines the If-Match condition. The operation will be performed only if the ETag on the server
// matches this value.
// ifNoneMatch - defines the If-None-Match condition. The operation will be performed only if the ETag on the
// server does not match this value.
func (client IndexersClient) CreateOrUpdate(ctx context.Context, indexerName string, indexer Indexer, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (result Indexer, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IndexersClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: indexer,
			Constraints: []validation.Constraint{{Target: "indexer.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "indexer.DataSourceName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "indexer.TargetIndexName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "indexer.Schedule", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "indexer.Schedule.Interval", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("searchservice.IndexersClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, indexerName, indexer, clientRequestID, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IndexersClient) CreateOrUpdatePreparer(ctx context.Context, indexerName string, indexer Indexer, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	pathParameters := map[string]interface{}{
		"indexerName": autorest.Encode("path", indexerName),
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/indexers('{indexerName}')", pathParameters),
		autorest.WithJSON(indexer),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Prefer", "return=representation"))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IndexersClient) CreateOrUpdateResponder(resp *http.Response) (result Indexer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an indexer.
// Parameters:
// indexerName - the name of the indexer to delete.
// clientRequestID - the tracking ID sent with the request to help with debugging.
// ifMatch - defines the If-Match condition. The operation will be performed only if the ETag on the server
// matches this value.
// ifNoneMatch - defines the If-None-Match condition. The operation will be performed only if the ETag on the
// server does not match this value.
func (client IndexersClient) Delete(ctx context.Context, indexerName string, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IndexersClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, indexerName, clientRequestID, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IndexersClient) DeletePreparer(ctx context.Context, indexerName string, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	pathParameters := map[string]interface{}{
		"indexerName": autorest.Encode("path", indexerName),
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/indexers('{indexerName}')", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IndexersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves an indexer definition.
// Parameters:
// indexerName - the name of the indexer to retrieve.
// clientRequestID - the tracking ID sent with the request to help with debugging.
func (client IndexersClient) Get(ctx context.Context, indexerName string, clientRequestID *uuid.UUID) (result Indexer, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IndexersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, indexerName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client IndexersClient) GetPreparer(ctx context.Context, indexerName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	pathParameters := map[string]interface{}{
		"indexerName": autorest.Encode("path", indexerName),
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/indexers('{indexerName}')", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IndexersClient) GetResponder(resp *http.Response) (result Indexer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetStatus returns the current status and execution history of an indexer.
// Parameters:
// indexerName - the name of the indexer for which to retrieve status.
// clientRequestID - the tracking ID sent with the request to help with debugging.
func (client IndexersClient) GetStatus(ctx context.Context, indexerName string, clientRequestID *uuid.UUID) (result IndexerExecutionInfo, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IndexersClient.GetStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetStatusPreparer(ctx, indexerName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "GetStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "GetStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "GetStatus", resp, "Failure responding to request")
		return
	}

	return
}

// GetStatusPreparer prepares the GetStatus request.
func (client IndexersClient) GetStatusPreparer(ctx context.Context, indexerName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	pathParameters := map[string]interface{}{
		"indexerName": autorest.Encode("path", indexerName),
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/indexers('{indexerName}')/search.status", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetStatusSender sends the GetStatus request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) GetStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetStatusResponder handles the response to the GetStatus request. The method always
// closes the http.Response Body.
func (client IndexersClient) GetStatusResponder(resp *http.Response) (result IndexerExecutionInfo, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all indexers available for a search service.
// Parameters:
// selectParameter - selects which top-level properties of the indexers to retrieve. Specified as a
// comma-separated list of JSON property names, or '*' for all properties. The default is all properties.
// clientRequestID - the tracking ID sent with the request to help with debugging.
func (client IndexersClient) List(ctx context.Context, selectParameter string, clientRequestID *uuid.UUID) (result ListIndexersResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IndexersClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, selectParameter, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client IndexersClient) ListPreparer(ctx context.Context, selectParameter string, clientRequestID *uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPath("/indexers"),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IndexersClient) ListResponder(resp *http.Response) (result ListIndexersResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Reset resets the change tracking state associated with an indexer.
// Parameters:
// indexerName - the name of the indexer to reset.
// clientRequestID - the tracking ID sent with the request to help with debugging.
func (client IndexersClient) Reset(ctx context.Context, indexerName string, clientRequestID *uuid.UUID) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IndexersClient.Reset")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ResetPreparer(ctx, indexerName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Reset", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResetSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Reset", resp, "Failure sending request")
		return
	}

	result, err = client.ResetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Reset", resp, "Failure responding to request")
		return
	}

	return
}

// ResetPreparer prepares the Reset request.
func (client IndexersClient) ResetPreparer(ctx context.Context, indexerName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	pathParameters := map[string]interface{}{
		"indexerName": autorest.Encode("path", indexerName),
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/indexers('{indexerName}')/search.reset", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResetSender sends the Reset request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) ResetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ResetResponder handles the response to the Reset request. The method always
// closes the http.Response Body.
func (client IndexersClient) ResetResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Run runs an indexer on-demand.
// Parameters:
// indexerName - the name of the indexer to run.
// clientRequestID - the tracking ID sent with the request to help with debugging.
func (client IndexersClient) Run(ctx context.Context, indexerName string, clientRequestID *uuid.UUID) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IndexersClient.Run")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RunPreparer(ctx, indexerName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Run", nil, "Failure preparing request")
		return
	}

	resp, err := client.RunSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Run", resp, "Failure sending request")
		return
	}

	result, err = client.RunResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "searchservice.IndexersClient", "Run", resp, "Failure responding to request")
		return
	}

	return
}

// RunPreparer prepares the Run request.
func (client IndexersClient) RunPreparer(ctx context.Context, indexerName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{

		"searchDnsSuffix":   client.SearchDNSSuffix,
		"searchServiceName": client.SearchServiceName,
	}

	pathParameters := map[string]interface{}{
		"indexerName": autorest.Encode("path", indexerName),
	}

	const APIVersion = "2019-05-06"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{searchServiceName}.{searchDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/indexers('{indexerName}')/search.run", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RunSender sends the Run request. The method will close the
// http.Response Body if it receives an error.
func (client IndexersClient) RunSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RunResponder handles the response to the Run request. The method always
// closes the http.Response Body.
func (client IndexersClient) RunResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
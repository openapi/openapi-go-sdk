package client

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// NewClient creates a new Client instance with the given authentication token.
// The returned client can be used to make API requests.
func NewClient(token string) *Client {
	return &Client{
		authHeader: "Bearer " + token,
		httpClient: &http.Client{},
	}
}

// Client is a simple HTTP client used to make API requests.
type Client struct {
	authHeader string
	httpClient *http.Client
}

// Request sends an HTTP request to the specified endpoint.
// ctx: the context.Context to be used for the request.
// baseURL: the base URL of the API.
// method: the HTTP method to use (e.g. GET, POST, etc.).
// endpoint: the endpoint to send the request to.
// body: the body of the request.
// params: the query parameters to include in the request.
//
// Returns a string containing the formatted JSON response and an error if any.
func (c *Client) Request(
	ctx context.Context,
	method string,
	baseURL string,
	endpoint string,
	body io.Reader,
	params map[string]string) (string, error) {

	queryParams := url.Values{}
	for key, value := range params {
		queryParams.Set(key, value)
	}
	opts := &requestOptions{
		method:      method,
		baseURL:     baseURL,
		endpoint:    endpoint,
		authString:  c.authHeader,
		payload:     body,
		queryParams: queryParams,
	}
	return request(ctx, c.httpClient, opts)
}
func (c *Client) RequestBytes(
	ctx context.Context,
	method string,
	baseURL string,
	endpoint string,
	body io.Reader,
	params map[string]string) ([]byte,error){
		queryParams := url.Values{}
	for key, value := range params {
		queryParams.Set(key, value)
	}
	opts := &requestOptions{
		method:      method,
		baseURL:     baseURL,
		endpoint:    endpoint,
		authString:  c.authHeader,
		payload:     body,
		queryParams: queryParams,
	}
	return requestBytes(ctx, c.httpClient, opts)
		
	}

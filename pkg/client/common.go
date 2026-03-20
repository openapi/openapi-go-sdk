package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// httpClient is an interface that defines the Do method for making HTTP requests.
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// requestOptions contains the options for an HTTP request.
type requestOptions struct {
	method      string     // HTTP request method (e.g., "GET", "POST", etc.)
	baseURL     string     // Base URL for the HTTP request.
	endpoint    string     // Endpoint for the HTTP request.
	authString  string     // Authentication string for the HTTP request.
	payload     io.Reader  // Payload for the HTTP request (if any).
	queryParams url.Values // Query parameters for the HTTP request (if any).
}

// request makes an HTTP request with the specified options.
func request(
	ctx context.Context,
	client httpClient,
	opts *requestOptions) (string, error) {

	req, err := http.NewRequestWithContext(ctx, opts.method, opts.baseURL, opts.payload)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", opts.authString)
	req.Header.Add("Content-Type", `application/json`)
	req.URL.Path += opts.endpoint

	if opts.queryParams != nil {
		req.URL.RawQuery = opts.queryParams.Encode()
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var formattedJSON bytes.Buffer
	err = json.Indent(&formattedJSON, body, "", "\t")
	if err != nil {
		return "", err
	}
	
	return formattedJSON.String(), nil
}

func requestBytes(
	ctx context.Context,
	client httpClient,
	opts *requestOptions) ([]byte, error) {

	req, err := http.NewRequestWithContext(ctx, opts.method, opts.baseURL, opts.payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", opts.authString)
	req.Header.Add("Content-Type", `application/json`)
	req.URL.Path += opts.endpoint

	if opts.queryParams != nil {
		req.URL.RawQuery = opts.queryParams.Encode()
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var formattedJSON bytes.Buffer
	err = json.Indent(&formattedJSON, body, "", "\t")
	if err != nil {
		return nil, err
	}
	
	return formattedJSON.Bytes(), nil
}

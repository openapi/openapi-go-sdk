package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	oauthBaseURL     = "https://oauth.openapi.it"
	testOauthBaseURL = "https://test.oauth.openapi.it"
)

// NewOauthClient creates a new instance of OauthClient, which is used for making OAuth API requests.
// It takes a username string, an apikey string, and a test boolean as arguments.
// It returns a pointer to an OauthClient instance.
func NewOauthClient(username string, apikey string, test bool) *OauthClient {
	auth := []byte(username + ":" + apikey)
	authHeader := base64.StdEncoding.EncodeToString(auth)
	baseURL := oauthBaseURL
	if test {
		baseURL = testOauthBaseURL
	}
	return &OauthClient{
		baseURL:    baseURL,
		authHeader: "Basic " + authHeader,
		httpClient: &http.Client{},
	}
}

// OauthClient is a client used for making OAuth API requests.
type OauthClient struct {
	baseURL    string
	authHeader string
	httpClient *http.Client
}

// GetScopes sends an HTTP GET request to the /scopes endpoint of the OAuth API.
// It takes a context.Context and a boolean limit as arguments.
// It returns a string containing the JSON response body, and an error if any.
func (c *OauthClient) GetScopes(ctx context.Context, limit bool) (string, error) {
	queryParams := url.Values{}
	if limit {
		queryParams.Set("limit", "1")
	}
	opts := &requestOptions{
		method:      "GET",
		baseURL:     c.baseURL,
		endpoint:    "/scopes",
		authString:  c.authHeader,
		payload:     nil,
		queryParams: queryParams,
	}

	return request(ctx, c.httpClient, opts)
}

// CreateToken creates a new token for the specified scopes with the given time-to-live (TTL).
// The method accepts a slice of strings specifying the scopes to be granted to the token, and an integer representing
// the TTL in seconds. The method returns a string containing the token and an error if any.
func (c *OauthClient) CreateToken(ctx context.Context, scopes []string, ttl int) (string, error) {
	payload := struct {
		Scopes []string `json:"scopes"`
		TTL    int      `json:"ttl"`
	}{scopes, ttl}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		return "", err
	}
	opts := &requestOptions{
		method:      "POST",
		baseURL:     c.baseURL,
		endpoint:    "/token",
		authString:  c.authHeader,
		payload:     &buf,
		queryParams: nil,
	}

	return request(ctx, c.httpClient, opts)
}

// GetTokens returns a list of all the tokens that have the specified scope.
// The method accepts a string containing the scope for which tokens are to be retrieved. The method returns
// a string containing the tokens and an error if any.
func (c *OauthClient) GetTokens(ctx context.Context, scope string) (string, error) {
	queryParams := url.Values{}
	queryParams.Set("scope", scope)
	opts := &requestOptions{
		method:      "GET",
		baseURL:     c.baseURL,
		endpoint:    "/token",
		authString:  c.authHeader,
		payload:     nil,
		queryParams: queryParams,
	}

	return request(ctx, c.httpClient, opts)
}

// DeleteToken deletes the token with the specified ID.
// The method accepts a string containing the ID of the token to be deleted. The method returns a string and an error if any.
func (c *OauthClient) DeleteToken(ctx context.Context, id string) (string, error) {
	endpoint := "/token" + "/" + id
	opts := &requestOptions{
		method:      "DELETE",
		baseURL:     c.baseURL,
		endpoint:    endpoint,
		authString:  c.authHeader,
		payload:     nil,
		queryParams: nil,
	}

	return request(ctx, c.httpClient, opts)
}

// GetCounters retrieves the counters for the specified period and date.
// The method accepts two strings - one containing the period for which counters are to be retrieved, and the other containing
// the date for which counters are to be retrieved. The method returns a string containing the counters and an error if any.
func (c *OauthClient) GetCounters(ctx context.Context, period string, date string) (string, error) {
	endpoint := "/counters" + "/" + period + "/" + date
	opts := &requestOptions{
		method:      "GET",
		baseURL:     c.baseURL,
		endpoint:    endpoint,
		authString:  c.authHeader,
		payload:     nil,
		queryParams: nil,
	}

	return request(ctx, c.httpClient, opts)
}

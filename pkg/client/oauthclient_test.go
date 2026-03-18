package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewOauthClient(t *testing.T) {
	c := NewOauthClient("user", "apikey", false)
	if c == nil {
		t.Fatal("expected non-nil oauth client")
	}
	if c.baseURL != oauthBaseURL {
		t.Errorf("expected base URL %q, got %q", oauthBaseURL, c.baseURL)
	}
	if !strings.HasPrefix(c.authHeader, "Basic ") {
		t.Errorf("expected Basic auth header, got %q", c.authHeader)
	}
}

func TestNewOauthClientTest(t *testing.T) {
	c := NewOauthClient("user", "apikey", true)
	if c.baseURL != testOauthBaseURL {
		t.Errorf("expected test base URL %q, got %q", testOauthBaseURL, c.baseURL)
	}
}

func TestGetScopes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/scopes" {
			t.Errorf("expected /scopes, got %q", r.URL.Path)
		}
		w.Write([]byte(`{"scopes":[]}`))
	}))
	defer server.Close()

	c := &OauthClient{
		baseURL:    server.URL,
		authHeader: "Basic dXNlcjpha2V5",
		httpClient: &http.Client{},
	}
	resp, err := c.GetScopes(context.Background(), false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(resp, "scopes") {
		t.Errorf("unexpected response: %s", resp)
	}
}

func TestGetScopesWithLimit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("limit") != "1" {
			t.Errorf("expected limit=1, got %q", r.URL.Query().Get("limit"))
		}
		w.Write([]byte(`{"scopes":[]}`))
	}))
	defer server.Close()

	c := &OauthClient{
		baseURL:    server.URL,
		authHeader: "Basic dXNlcjpha2V5",
		httpClient: &http.Client{},
	}
	_, err := c.GetScopes(context.Background(), true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCreateToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/token" {
			t.Errorf("expected /token, got %q", r.URL.Path)
		}
		w.Write([]byte(`{"token":"abc123","scopes":["GET:test.example.com"]}`))
	}))
	defer server.Close()

	c := &OauthClient{
		baseURL:    server.URL,
		authHeader: "Basic dXNlcjpha2V5",
		httpClient: &http.Client{},
	}
	resp, err := c.CreateToken(context.Background(), []string{"GET:test.example.com"}, 3600)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(resp, "token") {
		t.Errorf("unexpected response: %s", resp)
	}
}

func TestGetTokens(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("scope") != "GET:test.example.com" {
			t.Errorf("expected scope param, got %q", r.URL.Query().Get("scope"))
		}
		w.Write([]byte(`{"tokens":[]}`))
	}))
	defer server.Close()

	c := &OauthClient{
		baseURL:    server.URL,
		authHeader: "Basic dXNlcjpha2V5",
		httpClient: &http.Client{},
	}
	resp, err := c.GetTokens(context.Background(), "GET:test.example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(resp, "tokens") {
		t.Errorf("unexpected response: %s", resp)
	}
}

func TestDeleteToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if !strings.HasSuffix(r.URL.Path, "/token-id") {
			t.Errorf("expected path ending in /token-id, got %q", r.URL.Path)
		}
		w.Write([]byte(`{"deleted":true}`))
	}))
	defer server.Close()

	c := &OauthClient{
		baseURL:    server.URL,
		authHeader: "Basic dXNlcjpha2V5",
		httpClient: &http.Client{},
	}
	resp, err := c.DeleteToken(context.Background(), "token-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(resp, "deleted") {
		t.Errorf("unexpected response: %s", resp)
	}
}

func TestGetCounters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/counters/daily/2024-01-01" {
			t.Errorf("unexpected path: %q", r.URL.Path)
		}
		w.Write([]byte(`{"counters":[]}`))
	}))
	defer server.Close()

	c := &OauthClient{
		baseURL:    server.URL,
		authHeader: "Basic dXNlcjpha2V5",
		httpClient: &http.Client{},
	}
	resp, err := c.GetCounters(context.Background(), "daily", "2024-01-01")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(resp, "counters") {
		t.Errorf("unexpected response: %s", resp)
	}
}
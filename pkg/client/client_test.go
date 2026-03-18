package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient("test-token")
	if c == nil {
		t.Fatal("expected non-nil client")
	}
	if c.authHeader != "Bearer test-token" {
		t.Errorf("expected auth header 'Bearer test-token', got %q", c.authHeader)
	}
}

func TestClientRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("wrong authorization header: %q", r.Header.Get("Authorization"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	c := NewClient("test-token")
	resp, err := c.Request(context.Background(), "GET", server.URL, "/test", nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(resp, "status") {
		t.Errorf("unexpected response: %s", resp)
	}
}

func TestClientRequestWithParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("key") != "value" {
			t.Errorf("expected query param key=value, got %v", r.URL.Query())
		}
		w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	c := NewClient("token")
	_, err := c.Request(context.Background(), "GET", server.URL, "/test", nil, map[string]string{"key": "value"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
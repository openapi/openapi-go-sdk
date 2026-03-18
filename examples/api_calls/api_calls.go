package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/openapi/openapi-go-sdk/pkg/client"
)

func main() {
	ctx := context.Background()

	apiClient := client.NewClient("<your_access_token>")

	// GET request with query parameters
	params := map[string]string{
		"denominazione": "altravia",
		"provincia":     "RM",
		"codice_ateco":  "6201",
	}
	result, err := apiClient.Request(ctx, "GET", "https://test.imprese.openapi.it", "/advance", nil, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GET API Response: %s\n", result)

	// POST request with payload
	payload := struct {
		Limit int `json:"limit"`
		Query struct {
			CountryCode string `json:"country_code"`
		} `json:"query"`
	}{
		Limit: 10,
		Query: struct {
			CountryCode string `json:"country_code"`
		}{CountryCode: "IT"},
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		log.Fatal(err)
	}
	result, err = apiClient.Request(ctx, "POST", "https://test.postontarget.com", "/fields/country", &buf, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("POST API Response: %s\n", result)
}
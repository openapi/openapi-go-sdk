<div align="center">
  <a href="https://openapi.com/">
    <img alt="Openapi SDK for Go" src=".github/assets/repo-header-a3.png" >
  </a>

  <h1>Openapi® client for Go</h1>
  <h4>The perfect starting point to integrate <a href="https://openapi.com/">Openapi®</a> within your Go project</h4>

[![Build Status](https://github.com/openapi/openapi-go-sdk/actions/workflows/go.yml/badge.svg)](https://github.com/openapi/openapi-go-sdk/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/openapi-it/openapi-cli-go)](https://goreportcard.com/report/github.com/openapi-it/openapi-cli-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/openapi-it/openapi-cli-go.svg)](https://pkg.go.dev/github.com/openapi-it/openapi-cli-go)
[![License](https://img.shields.io/github/license/openapi/openapi-go-sdk)](LICENSE)
<br>
[![Linux Foundation Member](https://img.shields.io/badge/Linux%20Foundation-Silver%20Member-003778?logo=linux-foundation&logoColor=white)](https://www.linuxfoundation.org/about/members)
</div>

## Overview

A minimal and agnostic Go SDK for Openapi, providing only the core HTTP primitives needed to interact with any Openapi service.

## Pre-requisites

Before using the Openapi Go Client, you will need an account at [Openapi](https://console.openapi.com/) and an API key to the sandbox and/or production environment.

## Features

- **Agnostic Design**: No API-specific classes, works with any Openapi service
- **Minimal Dependencies**: Only requires Go 1.19+
- **OAuth Support**: Built-in OAuth client for token management
- **HTTP Primitives**: GET, POST, PUT, DELETE and other HTTP methods
- **Clean Interface**: Simple and idiomatic Go API

## What you can do

With the Openapi Go Client, you can easily interact with a variety of services in the Openapi Marketplace. For example, you can:

- 📩 **Send SMS messages** with delivery reports and custom sender IDs
- 💸 **Process bills and payments** in real time via API
- 🧾 **Send electronic invoices** securely to the Italian Revenue Agency
- 📄 **Generate PDFs** from HTML content, including JavaScript rendering
- ✉️ **Manage certified emails** and legal communications via Italian Legalmail

For a complete list of all available services, check out the [Openapi Marketplace](https://console.openapi.com/) 🌐

## Installation

```bash
go get github.com/openapi-it/openapi-cli-go
```

## Usage

```go
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	client "github.com/openapi-it/openapi-cli-go/pkg/client"
)

func main() {
	ctx := context.Background()

	// Initialize the OAuth client on the sandbox environment
	oauthClient := client.NewOauthClient("<your_username>", "<your_apikey>", true)

	// Create a token for a list of scopes
	scopes := []string{
		"GET:test.imprese.openapi.it/advance",
		"POST:test.postontarget.com/fields/country",
	}
	ttl := 3600
	resp, err := oauthClient.CreateToken(ctx, scopes, ttl)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the token response
	tokenResponse := struct {
		Scopes []string `json:"scopes"`
		Token  string   `json:"token"`
	}{}
	_ = json.Unmarshal([]byte(resp), &tokenResponse)

	// Initialize the API client with the token
	apiClient := client.NewClient(tokenResponse.Token)

	// GET request with query params
	params := map[string]string{
		"denominazione": "altravia",
		"provincia":     "RM",
		"codice_ateco":  "6201",
	}
	_, err = apiClient.Request(ctx, "GET", "https://test.imprese.openapi.it", "/advance", nil, params)
	if err != nil {
		log.Fatal(err)
	}

	// POST request with a payload
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
	_, err = apiClient.Request(ctx, "POST", "https://test.postontarget.com", "/fields/country", &buf, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Delete the token when done
	_, err = oauthClient.DeleteToken(ctx, tokenResponse.Token)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Contributing

Contributions are always welcome! Whether you want to report bugs, suggest new features, improve documentation, or contribute code, your help is appreciated.

See [docs/contributing.md](docs/contributing.md) for detailed instructions on how to get started. Please make sure to follow this project's [docs/code-of-conduct.md](docs/code-of-conduct.md) to help maintain a welcoming and collaborative environment.

## Authors

Meet the project authors:

- L. Paderi ([@lpaderiAltravia](https://www.github.com/lpaderiAltravia))
- Openapi Team ([@openapi-it](https://github.com/openapi-it))

## Partners

Meet our partners using Openapi or contributing to this SDK:

- [Blank](https://www.blank.app/)
- [Credit Safe](https://www.creditsafe.com/)
- [Deliveroo](https://deliveroo.it/)
- [Gruppo MOL](https://molgroupitaly.it/it/)
- [Jakala](https://www.jakala.com/)
- [Octotelematics](https://www.octotelematics.com/)
- [OTOQI](https://otoqi.com/)
- [PWC](https://www.pwc.com/)
- [QOMODO S.R.L.](https://www.qomodo.me/)
- [SOUNDREEF S.P.A.](https://www.soundreef.com/)

## Our Commitments

We believe in open source and we act on that belief. We became Silver Members
of the Linux Foundation because we wanted to formally support the ecosystem
we build on every day. Open standards, open collaboration, and open governance
are part of how we work and how we think about software.

## License

This project is licensed under the [MIT License](LICENSE).

The MIT License is a permissive open-source license that allows you to freely use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the software, provided that the original copyright notice and this permission notice are included in all copies or substantial portions of the software.

In short, you are free to use this SDK in your personal, academic, or commercial projects, with minimal restrictions. The project is provided "as-is", without any warranty of any kind, either expressed or implied, including but not limited to the warranties of merchantability, fitness for a particular purpose, and non-infringement.

For more details, see the full license text at the [MIT License page](https://choosealicense.com/licenses/mit/).
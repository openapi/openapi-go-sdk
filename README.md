
# OpenApi IT Go Client 

This client is used to interact with the API found at [openapi.it](https://openapi.it/)

## Pre-requisites

Before using the OpenApi IT Go Client, you will need an account at [openapi.it](https://openapi.it/) and an API key to the sandbox and/or production environment

## Installation

You can install the OpenApi IT Go Client with the following command using go get:

```bash
go get github.com/openapi-it/openapi-cli-go
```
    
## Usage

```go
// main.go
package main

import (
	client "github.com/openapi-it/openapi-cli-go/pkg/client"
)

func main() {
	// Initialize the oauth client on the sandbox environment
	ctx := context.Background()
	oauthClient := client.NewOauthClient("<your_username>", "<your_apikey>", true)

	// Create a token for a list of scopes
	scopes := []string{
		"GET:test.cap.openapi.it/cerca_comuni",
		"POST:test.postontarget.com/fields/country",
	}
	ttl := 2592000
	resp, err := oauthClient.CreateToken(ctx, scopes, ttl) // returns the json as string
	if err != nil {
		log.Fatal(err)
	}

	// The string response can be parsed into a custom object
	tokenResponse := struct {
		Scopes []string `json:"scopes"`
		Token  string   `json:"token"`
	}{}
	_ = json.Unmarshal([]byte(resp), &tokenResponse)

	// Initialize the client
	client := client.NewClient(tokenResponse.Token)

	// Make a request with params
	params := map[string]string{
		"denominazione": "altravia",
		"provincia":     "RM",
		"codice_ateco":  "6201",
	}
	_, err = client.Request(
		ctx,
		"GET",
		"https://test.imprese.openapi.it",
		"/advance",
		nil, params,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Make a request with a payload
	payload := struct {
		Limit int `json:"limit"`
		Query struct {
			CountryCode string `json:"country_code"`
		} `json:"query"`
	}{
		Limit: 0,
		Query: struct {
			CountryCode string `json:"country_code"`
		}{CountryCode: "IT"},
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		log.Fatal(err)
	}
	_, err = client.Request(
		ctx,
		"POST",
		"https://test.postontarget.com",
		"/fields/country",
		&buf,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Delete the token
	_, err = oauthClient.DeleteToken(ctx, tokenResponse.Token)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.

Please adhere to this project's `code of conduct`.


## License

[MIT](https://choosealicense.com/licenses/mit/)


## Authors

- [@maiku1008](https://www.github.com/maiku1008)
- [@openapi-it](https://github.com/openapi-it)

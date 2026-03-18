package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/openapi/openapi-go-sdk/pkg/client"
)

func main() {
	ctx := context.Background()

	// Initialize the OAuth client on the sandbox environment (test=true)
	oauthClient := client.NewOauthClient("<your_username>", "<your_apikey>", true)

	scopes := []string{
		"GET:test.imprese.openapi.it/advance",
		"POST:test.postontarget.com/fields/country",
	}
	ttl := 3600

	resp, err := oauthClient.CreateToken(ctx, scopes, ttl)
	if err != nil {
		log.Fatal(err)
	}

	tokenResponse := struct {
		Scopes []string `json:"scopes"`
		Token  string   `json:"token"`
	}{}
	if err := json.Unmarshal([]byte(resp), &tokenResponse); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated token: %s\n", tokenResponse.Token)
	fmt.Println("Token created successfully!")
}
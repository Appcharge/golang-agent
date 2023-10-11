package auth_method

import (
	"encoding/json"
	"fmt"
	"golang-agent/pkg/auth/config"
	"net/http"
	_ "strings"
)

type GoogleAuth interface {
	Authenticate(appID, token string) config.AuthenticationResult
}

type GoogleAuthInit struct{}

// Authenticate authenticates a user using Google authentication.
func (g *GoogleAuthInit) Authenticate(appID, token string) config.AuthenticationResult {
	url := "https://oauth2.googleapis.com/tokeninfo?access_token=" + token

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return config.AuthenticationResult{IsValid: false}
	}

	resp, err := client.Do(req)
	fmt.Println(resp)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return config.AuthenticationResult{IsValid: false}
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			fmt.Println("Error decoding response body:", err)
			return config.AuthenticationResult{IsValid: false}
		}

		audience, ok := data["aud"]
		if !ok {
			fmt.Println("Invalid response format:", data)
			return config.AuthenticationResult{IsValid: false}
		}

		if audience == appID {
			return config.AuthenticationResult{IsValid: true, UserID: data["sub"].(string)}
		}
	}

	return config.AuthenticationResult{IsValid: false}
}

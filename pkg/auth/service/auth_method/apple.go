package auth_method

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang-agent/pkg/auth/config"
	"net/http"
	"strings"
)

type AppleAuth interface {
	Authenticate(appID, token, appleSecretAPI string) config.AuthenticationResult
}

type AppleAuthInit struct{}

func (a *AppleAuthInit) Authenticate(appID, token, appleSecretAPI string) config.AuthenticationResult {
	if appleSecretAPI == "" {
		fmt.Println("Apple secret API is not provided.")
		return config.AuthenticationResult{IsValid: false}
	}

	url := "https://appleid.apple.com/auth/token"

	client := &http.Client{}
	requestBody := strings.NewReader(
		"client_id=" + appID +
			"&client_secret=" + appleSecretAPI +
			"&code=" + token +
			"&grant_type=authorization_code",
	)

	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return config.AuthenticationResult{IsValid: false}
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
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

		idToken, _ := data["id_token"].(string)

		userID := extractUserIDFromIDToken(idToken)

		return config.AuthenticationResult{IsValid: true, UserID: userID}
	}

	return config.AuthenticationResult{IsValid: false}
}

// extractUserIDFromIDToken extracts the user ID from the ID token.
func extractUserIDFromIDToken(idToken string) string {
	parts := strings.Split(idToken, ".")
	if len(parts) != 3 {
		return "" // Invalid ID token format
	}

	encodedPayload := parts[1]
	decodedPayload, err := base64.RawURLEncoding.DecodeString(encodedPayload)
	if err != nil {
		fmt.Println("Error decoding ID token payload:", err)
		return ""
	}

	var payload map[string]interface{}
	err = json.Unmarshal(decodedPayload, &payload)
	if err != nil {
		fmt.Println("Error unmarshalling ID token payload:", err)
		return ""
	}

	userID, _ := payload["sub"].(string)
	return userID
}

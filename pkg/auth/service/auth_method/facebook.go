package auth_method

import (
	"encoding/json"
	"fmt"
	"golang-agent/pkg/auth/config"
	"net/http"
	"net/url"
	_ "strings"
)

// FacebookAuth is an interface for Facebook authentication.
type FacebookAuth interface {
	Authenticate(appID, token, appSecret string) config.AuthenticationResult
}

type FacebookAuthInit struct{}

// Authenticate authenticates a user using Facebook authentication.
func (f *FacebookAuthInit) Authenticate(appID, token, appSecret string) config.AuthenticationResult {
	inputToken := url.QueryEscape(token)
	accessToken := url.QueryEscape(appID + "|" + appSecret)

	url := "https://graph.facebook.com/debug_token?input_token=" + inputToken + "&access_token=" + accessToken

	fmt.Println(url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return config.AuthenticationResult{IsValid: false}
	}

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

		dataMap, ok := data["data"].(map[string]interface{})
		if !ok {
			fmt.Println("Invalid response format:", data)
			return config.AuthenticationResult{IsValid: false}
		}

		isValid, _ := dataMap["is_valid"].(bool)
		userID, _ := dataMap["user_id"].(string)

		return config.AuthenticationResult{IsValid: isValid, UserID: userID}
	}

	return config.AuthenticationResult{IsValid: false}
}

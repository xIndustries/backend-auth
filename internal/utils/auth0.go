package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ValidateAuth0Token verifies the structure of an Auth0 token.
func ValidateAuth0Token(token string) error {
	if len(strings.Split(token, ".")) != 3 {
		return fmt.Errorf("invalid token format")
	}
	return nil
}

// GetAuth0UserInfo retrieves user information from the Auth0 Management API.
func GetAuth0UserInfo(auth0Domain, token string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://%s/userinfo", auth0Domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user info: %s", resp.Status)
	}

	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

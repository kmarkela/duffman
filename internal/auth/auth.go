package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Do(req *http.Request, auth *Auth) error {

	switch auth.Type {

	case "Oauth2":
		doOauth2(req, auth.Details)
	case "Bearer":
		doBearer(req, auth.Details["bearer"])
	}

	return nil
}

func doOauth2(req *http.Request, det map[string]string) error {

	data := url.Values{}
	data.Set("grant_type", det["client_credentials"])
	data.Set("client_id", det["clientID"])
	data.Set("client_secret", det["clientSecret"])
	data.Set("scope", det["scope"])

	client := &http.Client{}
	r, err := http.NewRequest("POST", det["accessTokenUrl"], strings.NewReader(data.Encode()))

	if err != nil {
		return fmt.Errorf("auth error. %s", err.Error())
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		return fmt.Errorf("auth error. %s", err.Error())
	}

	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("auth error. %s", err.Error())
	}

	// Parse the JSON response to get the access token
	var tokenResponse map[string]interface{}
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return fmt.Errorf("auth error. %s", err.Error())
	}
	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		return fmt.Errorf("auth error. %s", err.Error())
	}

	doBearer(req, accessToken)

	return nil

}

func doBearer(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}

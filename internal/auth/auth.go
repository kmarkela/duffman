package auth

import "net/http"

func Do(req *http.Request, auth *Auth) error {

	switch auth.Type {

	case "Oauth2":
		doOauth2()
	}

	return nil
}

func doOauth2() {}

func doBearer(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}

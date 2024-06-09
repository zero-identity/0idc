package oauth2

import (
	"net/http"

	"github.com/zero-identity/0idc/errors"
)

func ClientCredentialsFlow(r *http.Request) (response interface{}, status int) {
	var clientID string
	var clientSecret string
	var ok bool

	clientID, clientSecret, ok = r.BasicAuth()

	if !ok {
		clientID = r.FormValue("client_id")
		clientSecret = r.FormValue("client_secret")
	}

	client, ok := getClient(clientID)

	if !ok {
		status = 401
		response = errors.INVALID_CLIENT
		return
	}

	ok = verifyClientSecret(client, clientSecret)

	if !ok {
		status = 401
		response = errors.INVALID_CLIENT
		return
	}

	status = 200
	response = &SuccessfulResponse{
		AccessToken: "2YotnFZFEjr1zCsicMWpAA",
		TokenType:   "Bearer",
		ExpiresIn:   3600,
		Scope:       "resource.READ",
	}
	return
}

func getClient(id string) (string, bool) {
	if id != "s6BhdRkqt3" {
		return id, false
	}

	return id, true
}

func verifyClientSecret(clientID interface{}, secret string) bool {
	return clientID != nil && secret == "gX1fBat3bV"
}

package oauth2

import (
	"net/http"

	"github.com/zero-identity/0idc/errors"
)

const (
	CLIENT_CREDENTIALS = "client_credentials"
	GRANT_TYPE         = "grant_type"
)

func HandleTokenRequest(r *http.Request) (response interface{}, status int) {
	grantType := r.FormValue(GRANT_TYPE)

	if isEmpty(grantType) {
		status = 400
		response = errors.INVALID_REQUEST
		return
	}

	switch grantType {
	case CLIENT_CREDENTIALS: return ClientCredentialsFlow(r)
	default:
		status = 400
		response = errors.UNSUPPORTED_GRANT_TYPE
	}

	return
}

func isEmpty(s string) bool { return s == "" }

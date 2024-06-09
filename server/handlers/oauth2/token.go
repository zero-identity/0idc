package oauth2

import (
	"encoding/json"
	"net/http"

	"github.com/zero-identity/0idc/oauth2"
	"github.com/zero-identity/0idc/server"
)

var TokenHandler = func(w http.ResponseWriter, r *http.Request) {
	response, status := oauth2.HandleTokenRequest(r)
	bytes, _ := json.Marshal(response)
	writeCacheHeaders(w.Header())
	server.WriteResponse(w, status, bytes)
}

func writeCacheHeaders(headers http.Header) {
	headers.Set("Cache-Control", "no-store")
	headers.Set("Pragma", "no-cache")
}

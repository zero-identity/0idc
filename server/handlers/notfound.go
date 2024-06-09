package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/zero-identity/0idc/errors"
	"github.com/zero-identity/0idc/server"
)

const (
	status = 400
)

var NotFound = func(w http.ResponseWriter, r *http.Request) {
	log.Info().
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Str("status", strconv.Itoa(status)).
		Msg("resource not found")

	response, _ := json.Marshal(errors.INVALID_REQUEST)

	server.WriteResponse(w, status, response)
}

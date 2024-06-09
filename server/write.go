package server

import "net/http"

func WriteResponse(w http.ResponseWriter, status int, response []byte) {
	headers := w.Header()
	headers.Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	w.Write(response)
}

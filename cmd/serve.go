package cmd

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zero-identity/0idc/server/handlers"
	"github.com/zero-identity/0idc/server/handlers/oauth2"
)

func Serve() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	http.HandleFunc("POST /oauth2/token", oauth2.TokenHandler)
	http.HandleFunc("/", handlers.NotFound)
	err := http.ListenAndServe(":9096", nil)
	log.Fatal().Msg(err.Error())
}

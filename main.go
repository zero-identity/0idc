package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("hello world")
	// manager := manage.NewDefaultManager()
	// manager.SetClientTokenCfg(&manage.Config{
	// 	AccessTokenExp:    10 * time.Minute,
	// 	IsGenerateRefresh: false,
	// })

	// mongoConf := mongo.NewConfigNonReplicaSet(
	// 	"mongodb://127.0.0.1:27017",
	// 	"oauth2", // database name
	// 	"root",
	// 	"secret",
	// 	"serviceName",
	// )

	// // use mongodb token store
	// manager.MapTokenStorage(mongo.NewTokenStore(mongoConf))

	// // client memory store
	// clientStore := mongo.NewClientStore(mongoConf)
	// manager.MapClientStorage(clientStore)

	// clientStore.Create(&models.Client{
	// 	ID:     "clientid",
	// 	Secret: "secret",
	// 	Domain: "somedomain",
	// 	Public: false,
	// 	UserID: "test",
	// })

	// srv := server.NewDefaultServer(manager)
	// srv.SetAllowGetAccessRequest(true)
	// srv.SetClientInfoHandler(server.ClientFormHandler)

	// srv.UserAuthorizationHandler = func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	// 	return "clientid", nil
	// }

	// srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
	// 	log.Println("Internal Error:", err.Error())
	// 	return
	// })

	// srv.SetResponseErrorHandler(func(re *errors.Response) {
	// 	log.Println("Response Error:", re.Error.Error())
	// })

	// http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
	// 	err := srv.HandleAuthorizeRequest(w, r)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusBadRequest)
	// 	}
	// })

	// http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
	// 	srv.HandleTokenRequest(w, r)
	// })

	// log.Fatal(http.ListenAndServe(":9096", nil))
}

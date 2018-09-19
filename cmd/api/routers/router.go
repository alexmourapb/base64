package routers

import (

	api "github.com/alexmourapb/base64api/cmd/api/handlers"
	"github.com/gorilla/mux"
)



func Router(h *api.Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", api.HandleHealthCheck).Methods("GET")

	router.HandleFunc("/encode", api.HandleEncodeBase64).Methods("POST")
	router.HandleFunc("/decode", api.HandleDecodeBase64).Methods("POST")

	return router
}

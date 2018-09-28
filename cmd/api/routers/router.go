package routers

import (
	"github.com/gorilla/mux"

	apiHandlers "github.com/alexmourapb/base64api/cmd/api/handlers"
)

//Router define the routes to the resources and methods
func Router(h *apiHandlers.Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", apiHandlers.HandleHealthCheck).Methods("GET")

	router.HandleFunc("/encode", apiHandlers.HandleEncodeBase64).Methods("POST")
	router.HandleFunc("/decode", apiHandlers.HandleDecodeBase64).Methods("POST")

	return router
}

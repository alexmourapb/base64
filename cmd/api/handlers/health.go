package handlers

import (
	"fmt"
	"net/http"
)

//HandleHealthCheck provide a service to check if the service is available.
func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message":"live"}`)
}

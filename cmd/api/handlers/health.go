package handlers

import (
"fmt"
"net/http"
)

//HandleHealthCheck ...
func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message":"live"}`)
}

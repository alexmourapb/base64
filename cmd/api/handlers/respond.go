package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//Respond provide service to return personalized error
func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintln(w, &buf)
}

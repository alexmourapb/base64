package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/alexmourapb/base64api/internal/generator"
)

type pwdGeneratorResponse struct {
	Alert       string `json:"alert,omitempty"`
	NewPassword string `json:"new-password"`
}

func PwdRadomHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pwdSize := params["pwd_size"]
	size, err := strconv.Atoi(pwdSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var alert string
	if size == 0 {
		alert = "Using default value: 8"
		size = 8
	}

	restult := generate.GenerateRandom(size)

	response := pwdGeneratorResponse{
		Alert:       alert,
		NewPassword: restult,
	}

	Respond(w, r, http.StatusCreated, response)
}

func getReqID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(0).(string); ok {
		return reqID
	}
	return ""
}

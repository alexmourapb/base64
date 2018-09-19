package handlers

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	apiErrors "github.com/alexmourapb/base64api/internal/base"
)

type DecodeRequest struct {
	DecodeTXT string `json:"decodeTXT"`
}

func HandleDecodeBase64(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println("Error into ParserForm", err)
		Respond(w, r, http.StatusBadRequest, apiErrors.NewErrorMessageForParser())
		return
	}

	request := new(DecodeRequest)
	decoderJSON := json.NewDecoder(r.Body)
	if err := decoderJSON.Decode(&request); err != nil {
		errMsg := "Body should be a JSON object"
		errDetail := map[string]string{"Detail": err.Error()}
		log.Println("Error into Decoder", err)
		Respond(w, r, http.StatusBadRequest, apiErrors.NewErrorMessage(errMsg, errDetail))
		return
	}

	decodeBase, _ :=  base64.StdEncoding.DecodeString(request.DecodeTXT)

	Respond(w, r, http.StatusCreated, string(decodeBase))
}

package handlers

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	apiErrors "github.com/alexmourapb/base64api/internal/base"
)

type EncodeRequest struct {
	EncodeTXT string `json:"encodeTXT"`
}

func HandleEncodeBase64(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println("Error into ParserForm", err)
		Respond(w, r, http.StatusBadRequest, apiErrors.NewErrorMessageForParser())
		return
	}

	request := new(EncodeRequest)
	decoderJSON := json.NewDecoder(r.Body)
	if err := decoderJSON.Decode(&request); err != nil {
		errMsg := "Body should be a JSON object"
		errDetail := map[string]string{"Detail": err.Error()}
		log.Println("Error into Decoder", err)
		Respond(w, r, http.StatusBadRequest, apiErrors.NewErrorMessage(errMsg, errDetail))
		return
	}

	encodeBase := base64.StdEncoding.EncodeToString([]byte(request.EncodeTXT))

	Respond(w, r, http.StatusCreated, encodeBase)
}

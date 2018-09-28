package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	apiBase64 "github.com/alexmourapb/base64api/internal/base"
)

type EncodeRequest struct {
	EncodeTXT string `json:"encodeTXT"`
}

type EncodeResponse struct {
	EncodedTXT string `json:"encodedTXT"`
}

//HandleEncodeBase64 provide a service to encode text to base 64 code
func HandleEncodeBase64(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println("Error into ParserForm", err)
		Respond(w, r, http.StatusBadRequest, apiBase64.NewErrorMessageForParser())
		return
	}

	request := new(EncodeRequest)
	decoderJSON := json.NewDecoder(r.Body)
	if err := decoderJSON.Decode(&request); err != nil {
		errMsg := "Body should be a JSON object"
		errDetail := map[string]string{"Detail": err.Error()}
		log.Println("Error into Decoder", err)
		Respond(w, r, http.StatusBadRequest, apiBase64.NewErrorMessage(errMsg, errDetail))
		return
	}

	encodeBase := apiBase64.Encode64(request.EncodeTXT)

	response := new(EncodeResponse)
	response.EncodedTXT = encodeBase

	Respond(w, r, http.StatusCreated, response)
}

type DecodeRequest struct {
	DecodeTXT string `json:"decodeTXT"`
}

type DecodeResponse struct {
	DecodedTXT string `json:"decodedTXT"`
}

//HandleDecodeBase64 provide a service to decode text from base 64 code
func HandleDecodeBase64(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println("Error into ParserForm", err)
		Respond(w, r, http.StatusBadRequest, apiBase64.NewErrorMessageForParser())
		return
	}

	request := new(DecodeRequest)
	decoderJSON := json.NewDecoder(r.Body)
	if err := decoderJSON.Decode(&request); err != nil {
		errMsg := "Body should be a JSON object"
		errDetail := map[string]string{"Detail": err.Error()}
		log.Println("Error into Decoder", err)
		Respond(w, r, http.StatusBadRequest, apiBase64.NewErrorMessage(errMsg, errDetail))
		return
	}

	decodeBase, err := apiBase64.Decode64(request.DecodeTXT)
	if err != nil {
		errDetail := map[string]string{"Detail": err.Error()}
		Respond(w, r, http.StatusBadRequest, apiBase64.NewErrorMessage(decodeBase, errDetail))
		return
	}

	response := new(DecodeResponse)
	response.DecodedTXT = string(decodeBase)

	Respond(w, r, http.StatusCreated, response)
}
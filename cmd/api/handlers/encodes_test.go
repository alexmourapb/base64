package handlers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	apiHandlers "github.com/alexmourapb/base64api/cmd/api/handlers"
	apiRouters "github.com/alexmourapb/base64api/cmd/api/routers"
	apiBase64 "github.com/alexmourapb/base64api/internal/base"
)

//TestHandleEncodeBase64WithSuccess ...
func TestHandleEncodeBase64WithSuccess(t *testing.T) {

	postRequest := new(apiHandlers.EncodeRequest)
	postRequest.EncodeTXT = "test base64api"

	encodeLoad := new(bytes.Buffer)
	json.NewEncoder(encodeLoad).Encode(postRequest)

	testServerMock := apiBase64.MockHandlerEncodeWithSuccess()
	defer testServerMock.Close()
	os.Setenv("http://localhost:8181", testServerMock.URL)

	h := &apiHandlers.Handler{}

	r, err := http.NewRequest("POST", "/encode", encodeLoad)
	apiBase64.AssertOk(t, err)

	w := httptest.NewRecorder()

	apiRouters.Router(h).ServeHTTP(w, r)
	apiBase64.AssertEquals(t, http.StatusCreated, w.Code)

	body, errBody := ioutil.ReadAll(w.Body)
	apiBase64.AssertOk(t, errBody)

	response := &apiHandlers.EncodeResponse{}
	errJSON := json.Unmarshal(body, response)
	apiBase64.AssertOk(t, errJSON)
	apiBase64.AssertEquals(t, "dGVzdCBiYXNlNjRhcGk=", response.EncodedTXT)
}

//TestHandleDecodeBase64WithSuccess ...
func TestHandleDecodeBase64WithSuccess(t *testing.T) {

	postRequest := new(apiHandlers.DecodeRequest)
	postRequest.DecodeTXT = "dGVzdCBiYXNlNjRhcGk="

	decodeLoad := new(bytes.Buffer)
	json.NewEncoder(decodeLoad).Encode(postRequest)

	testServerMock := apiBase64.MockHandlerDecodeWithSuccess()
	defer testServerMock.Close()
	os.Setenv("http://localhost:8181", testServerMock.URL)

	h := &apiHandlers.Handler{}

	r, err := http.NewRequest("POST", "/decode", decodeLoad)
	apiBase64.AssertOk(t, err)

	w := httptest.NewRecorder()

	apiRouters.Router(h).ServeHTTP(w, r)
	apiBase64.AssertEquals(t, http.StatusCreated, w.Code)

	body, errBody := ioutil.ReadAll(w.Body)
	apiBase64.AssertOk(t, errBody)

	response := &apiHandlers.DecodeResponse{}
	errJSON := json.Unmarshal(body, response)
	apiBase64.AssertOk(t, errJSON)
	apiBase64.AssertEquals(t, "test base64api", response.DecodedTXT)
}

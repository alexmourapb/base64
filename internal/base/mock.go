package base

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

const (
	encodeURL    = "/encode"
	encodeMethod = "POST"
	decodeURL    = "/decode"
	decodeMethod = "POST"
)

//MockHandlerEncodeWithSuccess simulates the return of the encode request
func MockHandlerEncodeWithSuccess() *httptest.Server {
	testEncodeMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == encodeURL && r.Method == encodeMethod {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "%s", `{ "encodedTXT": "dGVzdCBiYXNlNjRhcGk=" }`)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "invalid")
	}))
	return testEncodeMock
}

//MockHandlerDecodeWithSuccess simulates the return of the decode request
func MockHandlerDecodeWithSuccess() *httptest.Server {
	testDecodeMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == decodeURL && r.Method == decodeMethod {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "%s", `{ "decodedTXT": "test base64api" }`)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "invalid")
	}))
	return testDecodeMock
}

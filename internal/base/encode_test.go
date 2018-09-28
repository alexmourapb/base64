package base_test

import (
	"testing"

	apiBase64 "github.com/alexmourapb/base64api/internal/base"
)

//TestEncodeWithSuccess ...
func TestEncodeWithSuccess(t *testing.T) {

	result := apiBase64.Encode64("test")
	apiBase64.AssertEquals(t, "dGVzdA==", result)

}

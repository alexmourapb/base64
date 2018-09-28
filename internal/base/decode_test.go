package base_test

import (
	"testing"

	apiBase64 "github.com/alexmourapb/base64api/internal/base"
)

//TestDecodeWithSuccess ...
func TestDecodeWithSuccess(t *testing.T) {

	result, err := apiBase64.Decode64("dGVzdA==")
	apiBase64.AssertOk(t, err)
	apiBase64.AssertEquals(t, "test", result)
}

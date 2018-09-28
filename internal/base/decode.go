package base

import (
	"encoding/base64"
)

//Decode64 decode text from base 64 code
func Decode64(decodeTXT string) (string, error) {
	response, err := base64.StdEncoding.DecodeString(decodeTXT)
	if err != nil {
		errMsg := "Invalid text to decode"
		return errMsg, err
	}
	return string(response), nil
}

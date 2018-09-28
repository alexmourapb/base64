package base

import "encoding/base64"

//Encode64 encode text to base 64 code
func Encode64(encodeTXT string) string {
	response := base64.StdEncoding.EncodeToString([]byte(encodeTXT))
	return string(response)
}

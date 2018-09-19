package base

//ErrorMessage ...
type ErrorMessage struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

//NewErrorMessage ...
func NewErrorMessage(message string, errors map[string]string) *ErrorMessage {
	errorMessage := &ErrorMessage{}
	errorMessage.Message = message
	errorMessage.Errors = errors

	return errorMessage
}

//NewErrorMessageForParser ...
func NewErrorMessageForParser() *ErrorMessage {
	return NewErrorMessage("Problems parsing JSON", nil)
}

//NewErrorMessageForDecoder ...
func NewErrorMessageForDecoder() *ErrorMessage {
	return NewErrorMessage("Body should be a JSON object", nil)
}

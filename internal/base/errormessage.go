package base

//ErrorMessage ...
type ErrorMessage struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

//NewErrorMessage returns detailed error
func NewErrorMessage(message string, errors map[string]string) *ErrorMessage {
	errorMessage := &ErrorMessage{}
	errorMessage.Message = message
	errorMessage.Errors = errors

	return errorMessage
}

//NewErrorMessageForParser returns custom error for Parser
func NewErrorMessageForParser() *ErrorMessage {
	return NewErrorMessage("Problems parsing JSON", nil)
}


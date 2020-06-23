package upcloud

import "fmt"

// Error represents an error response payload
type Error struct {
	Code    string `json:"error_code"`
	Message string `json:"error_message"`
}

// Error will return the error representation of the Error
// Note: This causes Error to match the error interface
func (e *Error) Error() string {
	// Return a formatted version of the error message and code
	return fmt.Sprintf("%s [%s]", e.Message, e.Code)
}

type errorResponse struct {
	Error *Error `json:"error"`
}

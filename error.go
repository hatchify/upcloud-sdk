package upcloud

import "fmt"

// Error represents an UpCloud API error response payload
type Error struct {
	// UpCloud error code
	Code string `json:"error_code"`
	// UpCloud error message
	Message string `json:"error_message"`
}

// Error will return the error representation of the Error
// Note: This causes Error to match the error interface
func (e *Error) Error() string {
	// Return a formatted version of the error message and code
	return fmt.Sprintf("%s [%s]", e.Message, e.Code)
}

// errorResponse is a response wrapper to match the UpCloud API payload
type errorResponse struct {
	Error *Error `json:"error"`
}

package upcloud

import "fmt"

// Error represents an error response payload
type Error struct {
	Code    string `json:"error_code"`
	Message string `json:"error_message"`
}

func (e *Error) Error() error {
	return fmt.Errorf("%s [%s]", e.Message, e.Code)
}

type errorResponse struct {
	Error *Error `json:"error"`
}

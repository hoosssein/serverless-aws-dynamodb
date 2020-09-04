package utils

import "fmt"

type HttpError struct {
	OriginalError error
	Code          int
	Message       string
}

func (h HttpError) UserError() string {
	if h.Message != "" {
		return h.Message
	} else if h.Code < 500 {
		return h.Error()
	}
	return "INTERNAL SERVER ERROR"
}

func (h HttpError) Error() string {
	return h.String()
}

func (h HttpError) String() string {
	return fmt.Sprintf("Error Code: %v, Error message: %v, originalError: %v", h.Code, h.Message, h.OriginalError)
}

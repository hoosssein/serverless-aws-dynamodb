package utils

type HttpError struct {
	Err  error
	Code int
}

func (h HttpError) Error() string {
	return h.Err.Error()
}

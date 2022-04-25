package models

type RequestError struct {
	Status  int
	Message error
}

func NewRequestError(status int, err error) *RequestError {
	return &RequestError{
		Status:  status,
		Message: err,
	}
}

func (e RequestError) Error() string {
	return e.Message.Error()
}

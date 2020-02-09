package model

type NegoError struct {
	Err     error  `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewNegoError(err error, message string, status int) *NegoError {
	return &NegoError{
		Err:     err,
		Message: message,
		Status:  status,
	}
}

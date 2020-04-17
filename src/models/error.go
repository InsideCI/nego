package models

//NegoError unifies errors between services and controllers.
type NegoError struct {
	Err     string `json:"error,omitempty"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//NewNegoError returns a new instance of NegoError.
func NewNegoError(message string, status int, err error) *NegoError {
	return &NegoError{
		Message: message,
		Status:  status,
		Err:     err.Error(),
	}
}

//NewNegoErrorConstant returns a new NegoError constant.
func NewNegoErrorConstant(message string, status int) *NegoError {
	return &NegoError{
		Message: message,
		Status:  status,
	}
}

package utils

import (
	"encoding/json"
	"net/http"

	"github.com/InsideCI/nego/src/models"
)

//Throw is a error handler function.
func Throw(w http.ResponseWriter, exception *models.NegoError, err error) {
	Error := models.NegoError{
		Message: exception.Message,
		Status:  exception.Status,
	}

	if err != nil {
		Error.Err = err.Error()
	}

	w.WriteHeader(exception.Status)
	_ = json.NewEncoder(w).Encode(Error)
}

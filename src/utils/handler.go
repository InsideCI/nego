package utils

import (
	"encoding/json"
	"net/http"

	"github.com/InsideCI/nego/src/models"
)

//Throw is a error handler function.
func Throw(w http.ResponseWriter, exception *models.NegoError, err error) {
	if err != nil {
		exception.Err = err.Error()
	}

	w.WriteHeader(exception.Status)
	_ = json.NewEncoder(w).Encode(exception)
}

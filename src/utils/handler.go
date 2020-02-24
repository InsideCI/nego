package utils

import (
	"encoding/json"
	"github.com/InsideCI/nego/src/model"
	"github.com/InsideCI/nego/src/utils/exceptions"
	"net/http"
)

func Throw(w http.ResponseWriter, exception *exceptions.NegoErrorConstant, err error) {
	Error := model.NegoError{
		Err:     err.Error(),
		Message: exception.Message,
		Status:  exception.Status,
	}
	w.WriteHeader(exception.Status)
	_ = json.NewEncoder(w).Encode(Error)
}

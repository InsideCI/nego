package exceptions

import "net/http"

type NegoErrorConstant struct {
	Message string
	Status  int
}

var (
	BadRequest        = &NegoErrorConstant{"Bad request", http.StatusBadRequest}
	InternalError     = &NegoErrorConstant{"Internal server error", http.StatusInternalServerError}
	InvalidAttributes = &NegoErrorConstant{"Invalid attributes, please check your payload", http.StatusBadRequest}
	NotFound          = &NegoErrorConstant{"Register not found", http.StatusNotFound}
	InvalidIdentifier = &NegoErrorConstant{"Invalid identifier value", http.StatusBadRequest}
)

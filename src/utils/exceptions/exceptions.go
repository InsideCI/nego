package exceptions

import (
	"net/http"

	"github.com/InsideCI/nego/src/models"
)

//NegoErrorConstant unifies errors between controllers and services.
var (
	BadRequest        = models.NewNegoErrorConstant("Bad request", http.StatusBadRequest)
	InternalError     = models.NewNegoErrorConstant("Internal server error", http.StatusInternalServerError)
	InvalidAttributes = models.NewNegoErrorConstant("Invalid attributes, please check your payload", http.StatusBadRequest)
	NotFound          = models.NewNegoErrorConstant("Register not found", http.StatusNotFound)
	InvalidIdentifier = models.NewNegoErrorConstant("Invalid identifier value", http.StatusBadRequest)
	NotAuthorized     = models.NewNegoErrorConstant("Not Authorized", http.StatusUnauthorized)
	NotRegistered     = models.NewNegoErrorConstant("User not registered", http.StatusBadRequest)
	WrongPassword     = models.NewNegoErrorConstant("Wrong password", http.StatusUnauthorized)
)

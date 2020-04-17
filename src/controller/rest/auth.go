package rest

import (
	"io"
	"net/http"
	"time"

	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/services"
	"github.com/InsideCI/nego/src/utils"
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/InsideCI/nego/src/utils/exceptions"
	"github.com/dgrijalva/jwt-go"
)

// AuthController is a controller that wraps a UserService.
type AuthController struct {
	db      *driver.DB
	service *services.UserService
}

// NewAuthController creates a new controller with a specific database engine.
func NewAuthController(db *driver.DB) *AuthController {
	return &AuthController{
		db:      db,
		service: services.NewUserService(),
	}
}

//Register registers an user if provided registration token is valid.
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("payload").(*models.User)
	if err := c.service.Register(c.db, user); err != nil {
		utils.Throw(w, exceptions.BadRequest, err)
		return
	}
}

//Login checks if user credentials are valid and returns a JWT.
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("payload").(*models.User)
	err := c.service.Login(c.db, user)
	if err != nil {
		utils.Throw(w, err, nil)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":   time.Now().Unix(),
	})

	tokenString, errToken := token.SignedString([]byte(constants.JwtTokenKey))
	if errToken != nil {
		utils.Throw(w, exceptions.InternalError, nil)
		return
	}

	io.WriteString(w, `{"token":"`+tokenString+`"}`)
}

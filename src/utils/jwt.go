package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/InsideCI/nego/src/utils/exceptions"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

//JWT abstracts necessary JWT fields.
type JWT struct {
	tokenClaim string
	tokenAuth  *jwtauth.JWTAuth
}

//NewJWT returns a new instance of a predefined settings for JWT middlewares.
func NewJWT() *JWT {
	return &JWT{
		tokenClaim: "email",
		tokenAuth:  jwtauth.New("HS256", []byte(constants.JwtTokenKey), nil),
	}
}

//Encode uses user data for JWT encoding.
func (j *JWT) Encode(email string) string {
	claims := jwt.MapClaims{
		j.tokenClaim: email,
		"exp":        time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":        time.Now().Unix(),
	}

	_, tokenString, _ := j.tokenAuth.Encode(claims)
	return tokenString
}

//Verifier finkds a token from a request.
func (j *JWT) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(j.tokenAuth)
}

//Decode returns a JWT claim.
func (j *JWT) Decode(r *http.Request) string {
	claim, _ := j.Authenticate(r)
	return claim
}

//Authenticate goes throught a JWT validation process.
func (j *JWT) Authenticate(r *http.Request) (string, error) {
	token, claims, err := jwtauth.FromContext(r.Context())
	if err != nil || token == nil {
		return "", errors.New("empty or invalid JWT")
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims[j.tokenClaim].(string), nil
}

//Authenticator is a middleware for handling JWT.
func (j *JWT) Authenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := j.Authenticate(r)
			if err != nil {
				Throw(w, exceptions.NotAuthorized, err)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

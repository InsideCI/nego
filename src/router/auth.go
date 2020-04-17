package router

import (
	"github.com/InsideCI/nego/src/controller/rest"
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/router/middlewares"
	"github.com/go-chi/chi"
)

//
func NewAuthRouter(db *driver.DB) func(r chi.Router) {
	generic := middlewares.GenericMiddleware{Type: models.User{}}
	handlers := rest.NewAuthController(db)

	return func(r chi.Router) {
		// Users
		r.With(generic.Persist).Post("/register", handlers.Register)
		r.With(generic.Persist).Post("/login", handlers.Login)

	}
}

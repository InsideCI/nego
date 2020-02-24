package router

import (
	"github.com/InsideCI/nego/src/controller/rest"
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/model"
	"github.com/InsideCI/nego/src/router/middlewares"
	"github.com/go-chi/chi"
)

// NewStudentRouter returns a new router for student endpoints.
func NewStudentRouter(db *driver.DB) func(r chi.Router) {
	generic := middlewares.GenericMiddleware{Type: model.Student{}}
	handlers := rest.NewStudentController(db)

	return func(r chi.Router) {
		// students
		r.With(generic.Payload).Post("/", handlers.Create)
		r.With(generic.Fetch).Get("/", handlers.Fetch)

		// students/{id}
		r.Route("/{id}", func(r chi.Router) {
			r.Use(generic.Identifier)
			r.Get("/", handlers.FetchOne)
			//Put
			//Delete
		})
	}
}

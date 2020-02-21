package router

import (
	"github.com/InsideCI/nego/src/controller/rest"
	"github.com/InsideCI/nego/src/driver"
	"github.com/go-chi/chi"
)

// NewTeacherRouter returns a new router for student endpoints.
func NewTeacherRouter(db *driver.DB) func(r chi.Router) {
	return func(r chi.Router) {

		handlers := rest.NewTeacherController(db)
		// students
		r.With(PayloadContext).Post("/", handlers.Create)
		//r.With(middlewares.QueryContext).Get("/", handlers.Fetch)

		// students/{id}
		//r.Route("/{id}", func(r chi.Router) {
		//	r.Use(middlewares.IDContext)
		//	r.Get("/", handlers.FetchOne)
		//	//Put
		//	//Delete
		//})

	}
}

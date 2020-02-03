package routers

import (
	"context"
	"encoding/json"
	"github.com/InsideCI/nego/src/controllers/rest"
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	router "github.com/InsideCI/nego/src/routers/middlewares"
	"github.com/go-chi/chi"
	"net/http"
)

func ValidStudent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// NewStudentRouter returns a new routers for student endpoints.
func NewStudentRouter(db *driver.DB) func(r chi.Router) {
	return func(r chi.Router) {
		handlers := rest.NewStudentController(db)

		// students
		r.With(CreateContext).With().Post("/", handlers.Create)
		r.Get("/", handlers.Fetch)

		// students/{id}
		r.Route("/{id}", func(r chi.Router) {
			r.Use(router.IDContext)
			r.Get("/", handlers.FetchOne)
			//Put
			//Delete
		})

	}
}

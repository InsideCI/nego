package router

import (
	"context"
	"github.com/InsideCI/nego/handler/api/student"
	"net/http"

	"github.com/InsideCI/nego/driver"
	"github.com/go-chi/chi"
)

func studentCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewStudentRouter returns a new router for student endpoints.
func NewStudentRouter(db *driver.DB) func(r chi.Router) {
	return func(r chi.Router) {
		StudentHandler := student.NewStudentHandler(db)
		r.Post("/", StudentHandler.Create)
		r.Get("/", StudentHandler.Fetch)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(studentCtx)
			r.Get("/", StudentHandler.FetchOne)
			//Put
			//Delete
		})

	}
}

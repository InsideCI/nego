package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/InsideCI/nego/driver"
	"github.com/InsideCI/nego/handler"
	"github.com/go-chi/chi"
)

func studentCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		registration := chi.URLParam(r, "registration")
		fmt.Println("Registration: ", registration)
		ctx := context.WithValue(r.Context(), "registration", registration)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewStudentRouter returns a new router for student endpoints.
func NewStudentRouter(db *driver.DB) func(r chi.Router) {
	return func(r chi.Router) {
		StudentHandler := handler.NewStudentHandler(db)
		r.Post("/", StudentHandler.Create)
		r.Get("/", StudentHandler.Fetch)
		r.Route("/{registration}", func(r chi.Router) {
			r.Use(studentCtx)
			r.Get("/", StudentHandler.FetchOne)
		})

	}
}

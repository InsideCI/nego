package router

import (
	"context"
	"encoding/json"
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/model"
	router "github.com/InsideCI/nego/src/router/generic"
	"github.com/InsideCI/nego/src/service/rest/student"
	"github.com/go-chi/chi"
	"net/http"
)

func CreateContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		}
		var students []model.Student
		if err := json.NewDecoder(r.Body).Decode(&students); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		for _, stud := range students {
			if stud.Matricula == 0 || stud.Nome == "" || stud.IDCurso == 0 {
				http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
				w.Header().Set("ERROR", "You can't send empty values for student.")
				return
			}
		}
		ctx := context.WithValue(r.Context(), "students", students)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ValidStudent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// NewStudentRouter returns a new router for student endpoints.
func NewStudentRouter(db *driver.DB) func(r chi.Router) {
	return func(r chi.Router) {
		handlers := student.NewStudentService(db)

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

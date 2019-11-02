package router

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/InsideCI/nego/handler/api/student"
	"github.com/InsideCI/nego/model"
	"net/http"
	"strconv"

	"github.com/InsideCI/nego/driver"
	"github.com/go-chi/chi"
)

func IDContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		_, err := strconv.ParseInt(fmt.Sprintf("%v", id), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CreateContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		}
		var students []model.Student
		if err := json.NewDecoder(r.Body).Decode(&students); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		for _, student := range students {
			if student.Matricula == 0 || student.Nome == "" || student.IDCurso == 0 {
				http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
				w.Header().Set("ERROR", "You can't send empty values for student.")
				return
			}
		}
		ctx := context.WithValue(r.Context(), "students", students)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewStudentRouter returns a new router for student endpoints.
func NewStudentRouter(db *driver.DB) func(r chi.Router) {
	return func(r chi.Router) {
		handlers := student.NewStudentHandler(db)

		// students/
		r.With(CreateContext).Post("/", handlers.Create)
		r.Get("/", handlers.Fetch)

		// students/{id}
		r.Route("/{id}", func(r chi.Router) {
			r.Use(IDContext)
			r.Get("/", handlers.FetchOne)
			//Put
			//Delete
		})

	}
}

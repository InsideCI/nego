package routers

import (
	"context"
	"encoding/json"
	"github.com/InsideCI/nego/src/controllers/rest"
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/routers/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

func ValidStudent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func PayloadContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		}
		var student models.Student
		if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		//TODO: CHANGE MANUAL VALID CHECK FOR validator PACKAGE
		if student.Matricula == 0 || student.Nome == "" || student.IDCurso == 0 {
			http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
			w.Header().Set("ERROR", "You can't send empty values for student.")
			return
		}

		ctx := context.WithValue(r.Context(), "student", student)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewStudentRouter returns a new routers for student endpoints.
func NewStudentRouter(db *driver.DB) func(r chi.Router) {
	return func(r chi.Router) {
		handlers := rest.NewStudentController(db)

		cors := cors.New(cors.Options{
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		})
		r.Use(cors.Handler)

		// students
		r.With(PayloadContext).Post("/", handlers.Create)
		r.With(middlewares.QueryContext).Get("/", handlers.Fetch)

		// students/{id}
		r.Route("/{id}", func(r chi.Router) {
			r.Use(middlewares.IDContext)
			r.Get("/", handlers.FetchOne)
			//Put
			//Delete
		})

	}
}

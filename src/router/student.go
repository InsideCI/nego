package router

import (
	"context"
	"encoding/json"
	"github.com/InsideCI/nego/src/controller/rest"
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/model"
	"github.com/InsideCI/nego/src/router/middlewares"
	"github.com/InsideCI/nego/src/utils"
	"github.com/InsideCI/nego/src/utils/exceptions"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func PayloadContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		}
		var student model.Student
		if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
			utils.Throw(w, exceptions.BadRequest, err)
			return
		}

		v := validator.New()
		if err := v.Struct(student); err != nil {
			utils.Throw(w, exceptions.InvalidAttributes, err)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", student)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewStudentRouter returns a new router for student endpoints.
func NewStudentRouter(db *driver.DB) func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(middlewares.Cors.Handler)

		handlers := rest.NewStudentController(db)
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

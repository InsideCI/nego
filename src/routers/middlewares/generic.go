package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/InsideCI/nego/src/models"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

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

func PayloadContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		}
		var students []models.Student
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

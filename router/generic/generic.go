package router

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
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

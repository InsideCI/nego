package middlewares

import (
	"context"
	"fmt"
	"github.com/InsideCI/nego/src/utils"
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/InsideCI/nego/src/utils/exceptions"
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
			utils.Throw(w, exceptions.InvalidIdentifier, err)
		}
		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func QueryContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys := r.URL.Query()

		params := make(map[string][]string)
		if len(keys) != 0 {
			for key, value := range keys {
				params[key] = value
			}
		} else {
			limit := strconv.Itoa(constants.MAXIMUM_FETCH)
			params["limit"] = append(params["limit"], limit)
		}

		ctx := context.WithValue(r.Context(), "params", params)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

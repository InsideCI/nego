package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/InsideCI/nego/src/utils"
	"github.com/InsideCI/nego/src/utils/constants"
	"github.com/InsideCI/nego/src/utils/exceptions"
	"github.com/go-chi/chi"
	"net/http"
	"reflect"
	"strconv"
)

type GenericMiddleware struct {
	Type interface{}
}

func (g *GenericMiddleware) output() interface{} {
	out := reflect.New(reflect.TypeOf(g.Type)).Interface()
	return out
}

func (g *GenericMiddleware) Identifier(next http.Handler) http.Handler {
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

func (g *GenericMiddleware) Fetch(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys := r.URL.Query()

		params := make(map[string][]string)

		//DEFAULT PARAMS
		params["limit"] = append(params["limit"], strconv.Itoa(constants.MAXIMUM_FETCH))

		if len(keys) != 0 {
			for key, value := range keys {
				params[key] = value
			}
		}

		ctx := context.WithValue(r.Context(), "params", params)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (g *GenericMiddleware) Payload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		}

		out := g.output()

		if err := json.NewDecoder(r.Body).Decode(out); err != nil {
			utils.Throw(w, exceptions.BadRequest, err)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", out)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/utils"
	"github.com/InsideCI/nego/src/utils/exceptions"
	"github.com/go-chi/chi"
	"github.com/mitchellh/mapstructure"
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

func (g *GenericMiddleware) exampleResolver(keys map[string][]string) interface{} {
	out := g.output()
	example := map[string]string{}

	for key, value := range keys {
		example[key] = value[len(value)-1]
	}

	cfg := &mapstructure.DecoderConfig{
		Result:  out,
		TagName: "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	decoder.Decode(example)

	return out
}

func (g *GenericMiddleware) paramsResolver(keys map[string][]string) interface{} {
	out := models.QueryParams{}

	if limit, ok := keys["limit"]; ok {
		out.Limit, _ = strconv.Atoi(limit[len(limit)-1])
	}
	if page, ok := keys["page"]; ok {
		out.Page, _ = strconv.Atoi(page[len(page)-1])
	}
	if sort, ok := keys["sort"]; ok {
		out.Order = sort
	}

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
		example := g.exampleResolver(keys)
		params := g.paramsResolver(keys)

		queryContext := context.WithValue(r.Context(), "params", params)
		ctx := context.WithValue(queryContext, "example", example)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (g *GenericMiddleware) Persist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		}

		payload := g.output()

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			utils.Throw(w, exceptions.BadRequest, err)
			return
		}

		var validation []reflect.Value

		f := reflect.ValueOf(payload).MethodByName("Valid")
		if f.IsValid() {
			validation = f.Call(validation)
		}

		//TODO: find a better way of handling validation value instead of using Sprintf.
		if len(validation) > 0 {
			val := reflect.ValueOf(validation[0]).Interface()
			err := fmt.Sprintf("%v", val)
			if err != "<nil>" {
				utils.Throw(w, exceptions.BadRequest, errors.New(err))
				return
			}
		}

		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

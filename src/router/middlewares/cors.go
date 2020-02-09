package middlewares

import "github.com/go-chi/cors"

var Cors = cors.New(cors.Options{
	AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: true,
	MaxAge:           300, // Maximum value not ignored by any of major browsers
})

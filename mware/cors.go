package mware

import (
	"net/http"

	"github.com/rs/cors"
)

func Cors(h http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "X-Auth-Token"},
		OptionsPassthrough: true,
		AllowCredentials:   true,
	})

	return c.Handler(h)
}

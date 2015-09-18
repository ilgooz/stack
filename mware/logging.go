package mware

import (
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
)

func Logging(h http.Handler) http.Handler {
	l := &negroni.Logger{log.New(os.Stdout, "[api] ", 0)}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nw := negroni.NewResponseWriter(w)
		l.ServeHTTP(nw, r, h.ServeHTTP)
	})
}

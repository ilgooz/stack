package mware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/ilgooz/cryptoutils"
)

func Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nw := negroni.NewResponseWriter(w)

		l := &negroni.Logger{
			log.New(os.Stdout, fmt.Sprintf("[api] request id: %s ", cryptoutils.RandToken(5)), 0),
		}

		l.ServeHTTP(nw, r, h.ServeHTTP)
	})
}

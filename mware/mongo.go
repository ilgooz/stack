package mware

import (
	"net/http"

	"github.com/ilgooz/stack/conf"
	"github.com/ilgooz/stack/ctx"
)

func Mongo(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := conf.M.Copy()
		defer db.Close()

		ctx.SetM(r, db)

		h.ServeHTTP(w, r)
	})
}

package mware

import (
	"net/http"

	"github.com/ilgooz/stack/ctx"
	"github.com/ilgooz/stack/model"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := ctx.CurrentUser(r)

		if user != nil {
			h.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
	})
}

func SetUser(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Auth-Token")

		if token == "" {
			ctx.SetCurrentUser(r, nil)
		} else {
			user, found, err := model.FindUserByToken(token, ctx.M(r))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if !found {
				ctx.SetCurrentUser(r, nil)
			} else {
				ctx.SetCurrentUser(r, &user)
			}
		}

		h.ServeHTTP(w, r)
	})
}

package mware

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/ilgooz/stack/model"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := model.CurrentUser(r)

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

		user, found, err := model.FindUserByToken(token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !found {
			context.Set(r, "user", nil)
		} else {
			context.Set(r, "user", &user)
		}

		h.ServeHTTP(w, r)
	})
}

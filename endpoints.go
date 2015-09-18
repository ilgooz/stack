package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilgooz/stack/mware"
	"github.com/ilgooz/stack/route"
	"github.com/justinas/alice"
)

var (
	appChain  = alice.New(mware.Logging, mware.Cors)
	authChain = appChain.Append(mware.Auth)
)

func handler() http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").Path("/users").Handler(authChain.ThenFunc(route.CreateUserHandler))
	r.Methods("GET").Path("/users").Handler(authChain.ThenFunc(route.ListUsersHandler))

	r.Methods("GET").Path("/version").Handler(appChain.ThenFunc(VersionHandler))

	// a dummy handler to log all the other requests that directs to not existent endpoints
	r.PathPrefix("/").Handler(appChain.Then(http.DefaultServeMux))

	return r
}

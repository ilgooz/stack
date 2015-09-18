package main

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/ilgooz/stack/mware"
	"github.com/ilgooz/stack/route"
	"github.com/justinas/alice"
)

var (
	appChain = alice.New(
		context.ClearHandler,
		mware.Logging,
		mware.Cors,
		mware.SetUser,
	)

	authChain = appChain.Append(mware.Auth)
)

func handler() http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").Path("/users").Handler(appChain.ThenFunc(route.CreateUserHandler))
	r.Methods("GET").Path("/users").Handler(appChain.ThenFunc(route.ListUsersHandler))
	r.Methods("GET").Path("/users/{id}").Handler(authChain.ThenFunc(route.GetUserHandler))
	r.Methods("GET").Path("/me").Handler(authChain.ThenFunc(route.GetMeHandler))

	r.Methods("POST").Path("/tokens").Handler(appChain.ThenFunc(route.CreateTokenHandler))

	r.Methods("GET").Path("/version").Handler(appChain.ThenFunc(VersionHandler))

	// a dummy handler to log all the other requests that directs to not existent endpoints
	r.PathPrefix("/").Handler(appChain.Then(http.DefaultServeMux))

	return r
}

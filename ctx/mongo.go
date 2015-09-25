package ctx

import (
	"net/http"

	"github.com/gorilla/context"

	"gopkg.in/mgo.v2"
)

func M(r *http.Request) *mgo.Session {
	return context.Get(r, "mongo").(*mgo.Session)
}

func SetM(r *http.Request, m *mgo.Session) {
	context.Set(r, "mongo", m)
}

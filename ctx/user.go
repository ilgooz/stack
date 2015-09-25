package ctx

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/ilgooz/stack/model"
)

func CurrentUser(r *http.Request) *model.User {
	if user := context.Get(r, "user"); user != nil {
		return user.(*model.User)
	}
	return nil
}

func SetCurrentUser(r *http.Request, u *model.User) {
	context.Set(r, "user", u)
}

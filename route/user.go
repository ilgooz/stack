package route

import (
	"net/http"

	"github.com/ilgooz/httpres"
	model "github.com/ilgooz/stack/model"
)

type UsersResponse struct {
	Users []model.User `json:"users"`
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	rp := UsersResponse{[]model.User{{"İlker"}, {"İbrahim"}}}

	httpres.Json(w, http.StatusOK, rp)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

}

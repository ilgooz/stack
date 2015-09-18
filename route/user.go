package route

import (
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/ilgooz/cryptoutils"
	"github.com/ilgooz/form"
	"github.com/ilgooz/httpres"
	"github.com/ilgooz/stack/conf"
	model "github.com/ilgooz/stack/model"
)

type UsersResponse struct {
	Users []model.User `json:"users"`
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []model.User{}

	if err := conf.MDB.C("users").Find(bson.M{}).All(&users); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpres.Json(w, http.StatusOK, UsersResponse{users})
}

type UserResponse struct {
	User model.User `json:"user"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	fields := CreateUserForm{}

	cef, err := form.Parse(&fields, w, r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if cef.HasError() {
		cef.Error.Send(http.StatusBadRequest)
		return
	}

	hash, err := cryptoutils.Hash(fields.Password, 5)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := model.User{
		ID:    bson.NewObjectId(),
		Name:  fields.Name,
		Email: strings.TrimSpace(fields.Email),
		Hash:  hash,
	}

	if err := conf.MDB.C("users").Insert(&user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpres.Json(w, http.StatusCreated, UserResponse{user})
}

type CreateUserForm struct {
	Name     string `form:"as:name"`
	Email    string `form:"as:email,email,required"`
	Password string `form:"as:password,min:3,required"`
}

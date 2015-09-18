package route

import (
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
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

	hash, err := cryptoutils.Hash(fields.Password, conf.PasswordLevel)
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
		if mgo.IsDup(err) {
			cef.Error.SendMessage("this email address already exists", http.StatusBadRequest)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := model.NewToken(user.ID)

	if err := conf.MDB.C("tokens").Insert(&token); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.AccessToken = token.Token

	httpres.Json(w, http.StatusCreated, UserResponse{user})
}

type CreateUserForm struct {
	Name     string `form:"as:name"`
	Email    string `form:"as:email,email,required"`
	Password string `form:"as:password,min:3,required"`
}

func GetMeHandler(w http.ResponseWriter, r *http.Request) {
	user := model.CurrentUser(r)
	httpres.Json(w, http.StatusOK, UserResponse{*user})
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user model.User

	if err := conf.MDB.C("users").FindId(bson.ObjectIdHex(id)).One(&user); err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		return
	}

	httpres.Json(w, http.StatusOK, UserResponse{user})
}

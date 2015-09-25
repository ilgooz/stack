package route

import (
	"log"
	"net/http"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/ilgooz/bsonutils"
	"github.com/ilgooz/cryptoutils"
	"github.com/ilgooz/form"
	"github.com/ilgooz/httpres"
	"github.com/ilgooz/paging"
	"github.com/ilgooz/stack/conf"
	"github.com/ilgooz/stack/ctx"
	model "github.com/ilgooz/stack/model"
)

type UsersResponse struct {
	CurrentPage     int          `json:"current_page"`
	TotalPagesCount int          `json:"total_pages_count"`
	Users           []model.User `json:"users"`
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	fields := ListUsersForm{}

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

	users := []model.User{}

	m := bson.M{}

	if fields.Name != "" {
		//todo: do full text search instead
		m["name"] = bson.RegEx{fields.Name, "i"}
	}

	q := ctx.M(r).DB("").C("users").Find(m)

	totalCount, err := q.Count()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	p := paging.Paging{
		Page:  fields.Page,
		Limit: fields.Limit,
		Count: totalCount,
	}.Calc()

	if err = q.
		Limit(p.Limit).
		Skip(p.Offset).
		Sort("-created_at").
		All(&users); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rp := UsersResponse{
		Users:           users,
		CurrentPage:     p.Page,
		TotalPagesCount: p.TotalPages,
	}

	httpres.Json(w, http.StatusOK, rp)
}

type ListUsersForm struct {
	Name  string `form:"as:name"`
	Page  int    `form:"as:page"`
	Limit int    `form:"as:limit"`
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
		ID:        bson.NewObjectId(),
		Name:      fields.Name,
		Email:     strings.TrimSpace(fields.Email),
		Hash:      hash,
		CreatedAt: time.Now(),
	}

	if err := ctx.M(r).DB("").C("users").Insert(&user); err != nil {
		if mgo.IsDup(err) {
			cef.Error.SendMessage("this email address already exists", http.StatusBadRequest)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := model.NewToken(user.ID)

	if err := ctx.M(r).DB("").C("tokens").Insert(&token); err != nil {
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
	id, ok := bsonutils.ObjectId(mux.Vars(r)["id"])

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var user model.User

	if err := ctx.M(r).DB("").C("users").FindId(id).One(&user); err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		return
	}

	httpres.Json(w, http.StatusOK, UserResponse{user})
}

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
	"github.com/ilgooz/eres"
	"github.com/ilgooz/formutils"
	"github.com/ilgooz/httpres"
	"github.com/ilgooz/paging"
	"github.com/ilgooz/stack/conf"
	"github.com/ilgooz/stack/ctx"
	model "github.com/ilgooz/stack/model"
)

type usersResponse struct {
	CurrentPage     int          `json:"current_page"`
	TotalPagesCount int          `json:"total_pages_count"`
	Users           []model.User `json:"users"`
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	fields := listUsersForm{}

	if formutils.ParseSend(w, r, &fields) {
		return
	}

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

	users := []model.User{}

	if err = q.
		Limit(p.Limit).
		Skip(p.Offset).
		Sort("-created_at").
		All(&users); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rp := usersResponse{
		Users:           users,
		CurrentPage:     p.Page,
		TotalPagesCount: p.TotalPages,
	}

	httpres.Json(w, http.StatusOK, rp)
}

type listUsersForm struct {
	Name  string `schema:"name"`
	Page  int    `schema:"page"`
	Limit int    `schema:"limit"`
}

type userResponse struct {
	User        model.User `json:"user"`
	AccessToken string     `json:"access_token,omitempty"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	fields := createUserForm{}

	if formutils.ParseSend(w, r, &fields) {
		return
	}

	hash, err := cryptoutils.Hash(fields.Password, conf.PasswordLevel)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := model.User{
		ID:           bson.NewObjectId(),
		Name:         fields.Name,
		Email:        strings.TrimSpace(fields.Email),
		PasswordHash: hash,
		CreatedAt:    time.Now(),
	}

	if err := ctx.M(r).DB("").C("users").Insert(&user); err != nil {
		if mgo.IsDup(err) {
			eres.New(w).AddField("email", "already exists").Send()
			return
		}

		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := model.NewToken(user.ID, false)

	if err := ctx.M(r).DB("").C("tokens").Insert(&token); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rp := userResponse{
		User:        user,
		AccessToken: token.Token,
	}

	httpres.Json(w, http.StatusCreated, rp)
}

type createUserForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email" validate:"email,required"`
	Password string `schema:"password" validate:"min=3,required"`
}

func GetMeHandler(w http.ResponseWriter, r *http.Request) {
	user := ctx.CurrentUser(r)
	httpres.Json(w, http.StatusOK, userResponse{User: *user})
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := bsonutils.ObjectId(mux.Vars(r)["id"])

	if err != nil {
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpres.Json(w, http.StatusOK, userResponse{User: user})
}

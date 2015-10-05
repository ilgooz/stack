package route

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/ilgooz/eres"
	"github.com/ilgooz/formutils"
	"github.com/ilgooz/httpres"
	"github.com/ilgooz/stack/ctx"
	"github.com/ilgooz/stack/model"
)

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}

func CreateTokenHandler(w http.ResponseWriter, r *http.Request) {
	fields := createTokenForm{}

	if formutils.ParseSend(w, r, &fields) {
		return
	}

	var user model.User

	if err := ctx.M(r).DB("").C("users").Find(bson.M{
		"email": strings.TrimSpace(fields.Email),
	}).One(&user); err != nil {
		if err == mgo.ErrNotFound {
			eres.New(w).SetMessage("bad credentials").Send()
			return
		}

		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(fields.Password)); err != nil {
		eres.New(w).SetMessage("bad credentials").Send()
		return
	}

	token := model.NewToken(user.ID)

	if err := ctx.M(r).DB("").C("tokens").Insert(&token); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpres.Json(w, http.StatusCreated, tokenResponse{token.Token})
}

type createTokenForm struct {
	Email    string `schema:"email" validate:"required"`
	Password string `schema:"password" validate:"required"`
}

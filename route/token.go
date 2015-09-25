package route

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/ilgooz/form"
	"github.com/ilgooz/httpres"
	"github.com/ilgooz/stack/ctx"
	"github.com/ilgooz/stack/model"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func CreateTokenHandler(w http.ResponseWriter, r *http.Request) {
	fields := CreateTokenForm{}

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

	var user model.User

	if err := ctx.M(r).DB("").C("users").Find(bson.M{
		"email": strings.TrimSpace(fields.Email),
	}).One(&user); err != nil {
		if err == mgo.ErrNotFound {
			cef.Error.SendMessage("bad credentials", http.StatusBadRequest)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(fields.Password)); err != nil {
		cef.Error.SendMessage("bad credentials", http.StatusBadRequest)
		return
	}

	token := model.NewToken(user.ID)
	if err := ctx.M(r).DB("").C("tokens").Insert(&token); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpres.Json(w, http.StatusCreated, TokenResponse{token.Token})
}

type CreateTokenForm struct {
	Email    string `form:"as:email,required"`
	Password string `form:"as:password,required"`
}

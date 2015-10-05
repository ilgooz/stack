package model

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Password  string        `json:"-" bson:"-"`
	Hash      string        `json:"-"`
	CreatedAt time.Time     `json:"-" bson:"created_at"`
}

func FindUserByToken(t string, m *mgo.Session) (User, bool, error) {
	var (
		user  User
		token Token
	)

	if err := m.DB("").C("tokens").Find(bson.M{"token": t}).One(&token); err != nil {
		if err == mgo.ErrNotFound {
			return user, false, nil
		}
		log.Println(err)
		return user, false, err
	}

	if err := m.DB("").C("users").FindId(token.UserID).One(&user); err != nil {
		if err == mgo.ErrNotFound {
			return user, false, nil
		}
		log.Println(err)
		return user, false, err
	}

	return user, true, nil
}

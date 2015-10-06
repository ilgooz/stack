package model

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	PasswordHash string        `json:"-" bson:"password_hash"`
	CreatedAt    time.Time     `json:"-" bson:"created_at"`
}

func FindUserByToken(t string, m *mgo.Session) (user User, found bool, err error) {
	var token Token

	if err := m.DB("").C("tokens").Find(bson.M{"token": t}).One(&token); err != nil {
		if err == mgo.ErrNotFound {
			return user, false, nil
		}

		return user, false, err
	}

	// reset token expire countdown
	if !token.UpdatedAt.IsZero() {
		if err = m.DB("").C("tokens").UpdateId(token.ID, bson.M{
			"$set": bson.M{"updated_at": time.Now()},
		}); err != nil {
			if err == mgo.ErrNotFound {
				return user, false, nil
			}

			return user, false, err
		}
	}

	if err := m.DB("").C("users").FindId(token.UserID).One(&user); err != nil {
		if err == mgo.ErrNotFound {
			return user, false, nil
		}

		return user, false, err
	}

	return user, true, nil
}

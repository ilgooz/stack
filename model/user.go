package modal

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Password string        `json:"-" bson:"-"`
	Hash     string        `json:"-"`
}

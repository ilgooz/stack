package conf

import (
	"log"

	"gopkg.in/mgo.v2"
)

func dialMongo() *mgo.Session {
	mongo, err := mgo.Dial(*mongoAddr)
	if err != nil {
		log.Fatalln(err)
	}
	return mongo
}

func ensureIndex() {
	s := M.Copy()
	defer s.Close()

	// users indexes
	uc := s.DB("").C("users")

	if err := uc.EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatalln(err)
	}

	// token index
	tc := s.DB("").C("tokens")

	if err := tc.EnsureIndex(mgo.Index{
		Key:         []string{"updated_at"},
		ExpireAfter: *tokenExpire,
	}); err != nil {
		log.Fatalln(err)
	}
}

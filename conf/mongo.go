package conf

import (
	"log"
	"time"

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

	uc := s.DB("").C("users")
	if err := uc.EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatalln(err)
	}

	tc := s.DB("").C("tokens")
	if err := tc.EnsureIndex(mgo.Index{
		Key:         []string{"created_at"},
		ExpireAfter: time.Hour * time.Duration(*tokenExpire),
	}); err != nil {
		log.Fatalln(err)
	}
}

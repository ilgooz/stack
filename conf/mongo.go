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

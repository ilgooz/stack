package conf

import (
	"flag"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var (
	addr        = flag.String("addr", ":3000", "Server Address")
	mongoAddr   = flag.String("mongo", "mongodb://127.0.0.1:27017/stack", "Mongodb Address")
	dbName      = flag.String("db", "stack", "Database Name")
	tokenExpire = flag.Duration("expire-token", 3*time.Hour, "Expire User Access Token After Hours")
	tokenSize   = flag.Int("token-size", 16, "Token Size")
)

var (
	Addr          string
	M             *mgo.Session
	PasswordLevel = 5
	TokenSize     int
)

func Load() {
	// advanced logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// parse commandline flags
	flag.Parse()

	// init configs for global access
	Addr = *addr
	TokenSize = *tokenSize
	M = dialMongo()
	ensureIndex()
}

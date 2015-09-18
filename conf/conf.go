package conf

import (
	"flag"
	"log"

	"gopkg.in/mgo.v2"
)

var (
	addr      = flag.String("addr", ":3000", "Server Address")
	mongoAddr = flag.String("mongo", ":3000", "Mongodb Address")
)

var (
	Addr  string
	Mongo *mgo.Session
)

func Load() {
	// advanced logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// parse commandline flags
	flag.Parse()

	// init configs for global access
	Addr = *addr

	// var err error
	// Mongo, err = mgo.Dial(*mongoAddr)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}

package main

import (
	"log"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/ilgooz/stack/conf"
)

func main() {
	conf.Load()
	run()
}

func run() {
	server := http.Server{
		Addr:    conf.Addr,
		Handler: handler(),
	}
	log.Fatalln(gracehttp.Serve(&server))
}

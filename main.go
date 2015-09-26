package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ilgooz/stack/conf"
	"github.com/tylerb/graceful"
)

func main() {
	conf.Load()
	run()
}

func run() {
	server := &http.Server{
		Addr:    conf.Addr,
		Handler: handler(),
	}

	gserver := &graceful.Server{
		Timeout: 10 * time.Second,
		Server:  server,
	}

	fmt.Printf("api server started at: %s\n", conf.Addr)
	log.Fatalln(gserver.ListenAndServe())
}

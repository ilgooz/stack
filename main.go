package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ilgooz/stack/conf"
)

func main() {
	conf.Load()
	run()
}

func run() {
	fmt.Println(fmt.Sprintf("server listening at %s", conf.Addr))
	log.Fatalln(http.ListenAndServe(conf.Addr, handler()))
}

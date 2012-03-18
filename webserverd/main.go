package main

import (
	"webserver/handlers"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/foo", handlers.FooHandler)
	http.HandleFunc("/bar", handlers.BarHandler)
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}


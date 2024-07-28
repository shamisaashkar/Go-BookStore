package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shamisaashkar/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))

}

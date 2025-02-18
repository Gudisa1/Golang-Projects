package main

import (
	"book-store/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	r:=mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}
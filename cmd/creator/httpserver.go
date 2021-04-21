package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/private/creator/httpadapter"
)

func main() {
	router := mux.NewRouter()
	server := httpadapter.NewServer(router)
	server.RegisterRoutes()

	log.Println("Listening...")
	http.ListenAndServe(":3000", router)
}

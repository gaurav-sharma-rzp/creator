package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/private/creator/api"
)

func main() {
	router := mux.NewRouter()
	server := api.NewServer(router)
	server.RegisterRoutes()

	log.Println("Listening...")
	http.ListenAndServe(":3000", router)
}

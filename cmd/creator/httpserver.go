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
	port := ":3000"
	log.Println("listening on port ", port)
	http.ListenAndServe(port, router)
}

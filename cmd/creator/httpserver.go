package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/private/creator/httpadapter"
)

func main() {
	router := mux.NewRouter()
	server := httpadapter.NewServer(router)
	server.RegisterRoutes()
	port := ":3000"
	log.Println("listening on port ", port)
	err := http.ListenAndServe(port, router)
	fmt.Println(err)
}

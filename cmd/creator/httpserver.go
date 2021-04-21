package main

import (
	"log"
	"net/http"
	"time"

	"github.com/private/creator/httpadapter"
)

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	router := http.NewServeMux()
	server := httpadapter.NewServer(router)
	server.RegisterRoutes()

	th := &timeHandler{format: time.RFC1123}
	router.Handle("/time", th)

	log.Println("Listening...")
	http.ListenAndServe(":3000", router)
}

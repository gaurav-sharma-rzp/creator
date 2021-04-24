package httpadapter

import (
	"fmt"
	"net/http"
)

func (s *server) addHealthRoutes() {
	s.Router.HandleFunc("/status", handleStatus()).Methods(http.MethodGet)
}

func handleStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"%s\" : \"%s\"}", "status", "ok")
	}
}

package httpadapter

import (
	"fmt"
	"net/http"
)

func (s *server) addFloraRoutes() {
	s.Router.HandleFunc("flora/create", handleFloraCreate()).Methods("GET")
}

func handleFloraCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", "CREATED")
	}
}

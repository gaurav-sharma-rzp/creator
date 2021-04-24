package httpadapter

import (
	"github.com/gorilla/mux"
)

type server struct {
	Router *mux.Router
}

func NewServer(router *mux.Router) *server {
	return &server{
		Router: router,
	}
}

func (s *server) RegisterRoutes() {
	s.addHealthRoutes()
	s.addFloraRoutes()
	s.addLSystemRoutes()
}

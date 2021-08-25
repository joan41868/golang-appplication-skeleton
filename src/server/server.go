package server

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Server struct {
	handler  *mux.Router
	listener net.Listener
	port     string
}

func NewServer() *Server {
	ln, err := net.Listen("tcp", ":http")
	if err != nil {
		log.Println(err)
	}
	return &Server{
		handler:  mux.NewRouter(),
		listener: ln,
		port:     os.Getenv("PORT"),
	}
}


func (server *Server) AddRoute(path, method string, handlerFn func(wr http.ResponseWriter, r *http.Request)) {
	server.handler.HandleFunc(path, handlerFn).Methods(method)
}

func (server Server) Start() {
	http.ListenAndServe(server.port, server.handler)
}

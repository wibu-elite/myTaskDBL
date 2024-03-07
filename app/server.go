package app

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {
	fmt.Println("Welcome to GoToko")
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to Port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	var server = Server{}

	server.Initialize()
	server.Run(":9000")

}

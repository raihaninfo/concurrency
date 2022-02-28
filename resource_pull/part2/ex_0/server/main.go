package main

import (
	"fmt"
	"net/http"
	"time"
)

type Server interface {
	Start() error
	Stop()
}

type TCPServer struct {
	port string
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
	var srv Server = &TCPServer{port: "8081"}
	srv.Start()
	time.Sleep(10 * time.Second)
	srv.Stop()
}

func (srv *TCPServer) Start() error {
	fmt.Printf("Starting server on port %v", srv.port)
	return nil
}

func (srv *TCPServer) Stop() {
	fmt.Printf("Stopping server on port %v", srv.port)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello from http handler")
}

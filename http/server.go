package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HttpServer struct {
	Host   string
	Port   string
	Server http.Server
}

func NewHttpServer(host, port string) *HttpServer {
	return &HttpServer{
		Host: host,
		Port: port,
	}
}

func (s HttpServer) Listen() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetCategoryTree)

	s.Server = http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: mux,
	}

	log.Printf("server listening at %v", s.Server.Addr)
	fmt.Println(s.Server.ListenAndServe())
}

func (s HttpServer) Shutdown() {
	s.Server.Shutdown(context.Background())
}

func GetCategoryTree(w http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(w).Encode(ExampleResponse)
	if err != nil {
		panic(err)
	}
}

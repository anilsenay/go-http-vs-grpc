package main

import (
	config "github.com/anilsenay/go-http-vs-grpc"
	"github.com/anilsenay/go-http-vs-grpc/http"
)

func main() {
	server := http.NewHttpServer(config.HTTP_HOST, config.HTTP_PORT)
	server.Listen()
}

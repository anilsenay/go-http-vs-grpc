package main

import (
	config "http-vs-grpc"
	"http-vs-grpc/http"
)

func main() {
	server := http.NewHttpServer(config.HTTP_HOST, config.HTTP_PORT)
	server.Listen()
}

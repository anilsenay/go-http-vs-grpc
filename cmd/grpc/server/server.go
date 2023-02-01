package main

import (
	config "github.com/anilsenay/go-http-vs-grpc"
	r "github.com/anilsenay/go-http-vs-grpc/grpc"
)

func main() {
	server := r.NewGrpcServer(config.RPC_HOST, config.RPC_PORT)
	server.Serve()
}

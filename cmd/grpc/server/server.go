package main

import (
	config "http-vs-grpc"
	r "http-vs-grpc/grpc"
)

func main() {
	server := r.NewGrpcServer(config.RPC_HOST, config.RPC_PORT)
	server.Serve()
}

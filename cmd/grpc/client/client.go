package main

import (
	config "github.com/anilsenay/go-http-vs-grpc"
	r "github.com/anilsenay/go-http-vs-grpc/grpc"
)

func main() {
	client := r.NewGrpcClient(config.RPC_HOST + ":" + config.RPC_PORT)
	client.GetCategoryTree()
}

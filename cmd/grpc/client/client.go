package main

import (
	config "http-vs-grpc"
	r "http-vs-grpc/grpc"
)

func main() {
	client := r.NewGrpcClient(config.RPC_HOST + ":" + config.RPC_PORT)
	client.GetCategoryTree()
}

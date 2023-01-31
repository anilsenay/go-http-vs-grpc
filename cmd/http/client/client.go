package main

import (
	config "http-vs-grpc"
	"http-vs-grpc/http"
)

func main() {
	client := http.NewHttpClient("http://" + config.HTTP_HOST + ":" + config.HTTP_PORT + "/")
	client.SendRequest()
}

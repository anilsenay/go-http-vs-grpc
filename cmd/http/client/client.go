package main

import (
	config "github.com/anilsenay/go-http-vs-grpc"
	"github.com/anilsenay/go-http-vs-grpc/http"
)

func main() {
	client := http.NewHttpClient("http://" + config.HTTP_HOST + ":" + config.HTTP_PORT + "/")
	client.SendRequest()
}

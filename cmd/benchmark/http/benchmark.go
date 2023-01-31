package main

import (
	"fmt"
	config "http-vs-grpc"
	r "http-vs-grpc/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	http_server := r.NewHttpServer(config.RPC_HOST, config.RPC_PORT)
	http_client := r.NewHttpClient("http://" + config.RPC_HOST + ":" + config.RPC_PORT)

	go http_server.Listen()

	argsWithoutProg := os.Args[1:]

	var requestNumber int
	var concurrency int = 1

	r, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		panic(err)
	}
	requestNumber = r

	if len(argsWithoutProg) >= 2 {
		c, err := strconv.Atoi(argsWithoutProg[1])
		if err != nil {
			panic(err)
		}
		concurrency = c
	}

	var durationPerRequest = make([]time.Duration, requestNumber)

	wg := sync.WaitGroup{}
	wg.Add(requestNumber)
	var ch = make(chan struct{}, concurrency)

	for i := 0; i < requestNumber; i++ {
		ch <- struct{}{}
		go func(i int) {
			start := time.Now()
			http_client.SendRequest()
			durationPerRequest[i] = time.Since(start)
			<-ch
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(durationPerRequest)
}

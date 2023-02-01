package benchmark

import (
	"sync"
	"time"

	config "github.com/anilsenay/go-http-vs-grpc"
	r "github.com/anilsenay/go-http-vs-grpc/http"
)

func RunBenchmark(requestNumber, concurrency int) []time.Duration {
	http_server := r.NewHttpServer(config.HTTP_HOST, config.HTTP_PORT)
	http_client := r.NewHttpClient("http://" + config.HTTP_HOST + ":" + config.HTTP_PORT)

	go http_server.Listen()

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

	return durationPerRequest
}

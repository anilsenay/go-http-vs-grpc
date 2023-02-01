package benchmark

import (
	"sync"
	"time"

	config "github.com/anilsenay/go-http-vs-grpc"
	r "github.com/anilsenay/go-http-vs-grpc/grpc"
)

func RunBenchmark(requestNumber, concurrency int) []time.Duration {
	grpc_server := r.NewGrpcServer(config.RPC_HOST, config.RPC_PORT)
	grpc_client := r.NewGrpcClient(config.RPC_HOST + ":" + config.RPC_PORT)

	go grpc_server.Serve()

	var durationPerRequest = make([]time.Duration, requestNumber)

	wg := sync.WaitGroup{}
	wg.Add(requestNumber)
	var ch = make(chan struct{}, concurrency)

	for i := 0; i < requestNumber; i++ {
		ch <- struct{}{}
		go func(i int) {
			start := time.Now()
			grpc_client.GetCategoryTree()
			durationPerRequest[i] = time.Since(start)
			<-ch
			wg.Done()
		}(i)
	}

	wg.Wait()

	return durationPerRequest
}

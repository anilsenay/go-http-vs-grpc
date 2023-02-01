package httpvsgrpc

import (
	"fmt"
	"testing"
	"time"

	r "github.com/anilsenay/go-http-vs-grpc/grpc"
	"github.com/anilsenay/go-http-vs-grpc/http"
)

// Run via:
// go test -bench=. -benchtime=1000x -benchmem

func BenchmarkServer(b *testing.B) {
	fmt.Println("================== BenchmarkServer ==================")

	http_server := http.NewHttpServer(HTTP_HOST, HTTP_PORT)
	http_client := http.NewHttpClient("http://" + HTTP_HOST + ":" + HTTP_PORT + "/")

	grpc_server := r.NewGrpcServer(RPC_HOST, RPC_PORT)
	grpc_client := r.NewGrpcClient(RPC_HOST + ":" + RPC_PORT)

	go grpc_server.Serve()
	go http_server.Listen()

	var table = []struct {
		name string
		f    func()
	}{
		{name: "http", f: http_client.SendRequest},
		{name: "grpc", f: grpc_client.GetCategoryTree},
	}

	for _, v := range table {
		b.Run(fmt.Sprintf("server_%s", v.name), func(b *testing.B) {
			start := time.Now()
			for i := 0; i < b.N; i++ {
				v.f()
			}
			fmt.Printf("\nserver_%s: took %d ns \n", v.name, time.Since(start))
		})
	}

	fmt.Println("closing...")
	http_server.Shutdown()
	grpc_server.Shutdown()

}

func BenchmarkServer_NewClientEachRequest(b *testing.B) {
	fmt.Println("================== BenchmarkServer_NewClientEachRequest ==================")

	http_server := http.NewHttpServer(HTTP_HOST, "11111")

	grpc_server := r.NewGrpcServer(RPC_HOST, "12111")

	go grpc_server.Serve()
	go http_server.Listen()

	b.Run("server_http", func(b *testing.B) {
		start := time.Now()
		for i := 0; i < b.N; i++ {
			http_client := http.NewHttpClient("http://" + HTTP_HOST + ":" + HTTP_PORT + "/")
			http_client.SendRequest()
		}
		fmt.Printf("\nserver_http: took %d ns \n", time.Since(start))
	})

	b.Run("server_grpc", func(b *testing.B) {
		start := time.Now()
		for i := 0; i < b.N; i++ {
			grpc_client := r.NewGrpcClient(RPC_HOST + ":" + RPC_PORT)
			grpc_client.GetCategoryTree()
		}
		fmt.Printf("\nserver_grpc: took %d ns \n", time.Since(start))
	})

	fmt.Println("closing...")
}

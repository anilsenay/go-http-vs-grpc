### Notes:

- I tested sending the same struct in both gRPC and HTTP and made simple benchmarks recording their response times.
- There are some benchmark tests in `request_benchmark_test.go` file. You can run them with `go test -bench=. -benchtime=1000x -benchmem`. There are 2 tests as `BenchmarkServer` and `BenchmarkServer_NewClientEachRequest`. I tested that what if it creates new gRPC connection for each request or use a single one for all requests. According to the results of those benchmarks, creating new connection for each request is not efficient for gRPC while it loses advantages of itself.
- Results may change by runs but gRPC was always better in my results.
- This test may not accurately reflect reality, or they may have wrong implementations etc. If you think there are some issues you can open an issue or contribute directly.

## Run benchmarks

#### Generate plots:

> Draws plots for grpc and http benchmarks

usage: `go run  cmd/benchmark/plots/plots.go <number of requests> <number of concurrent request>`

```bash
go run cmd/benchmark/plots/plots.go 1000 20
```

Last run results:

![echarts (2)](https://user-images.githubusercontent.com/1047345/215920137-c22e4091-53a8-4c79-b7cf-4198c713a918.png)


#### gRPC benchmark:

> Prints time values in array

usage: ` cmd/benchmark/grpc/main.go <number of requests> <number of concurrent request>`

```bash
go run cmd/benchmark/grpc/main.go 1000 20
```

#### HTTP benchmark:

> Prints time values in array

usage: ` cmd/benchmark/http/main.go <number of requests> <number of concurrent request>`

```bash
go run cmd/benchmark/http/main.go 1000 20
```

## Run server/client

#### gRPC server:

```bash
go run cmd/grpc/server/server.go
```

#### gRPC client:

```bash
go run cmd/grpc/client/client.go
```

#### HTTP server:

```bash
go run cmd/http/server/server.go
```

#### HTTP client:

```bash
go run cmd/http/client/client.go
```

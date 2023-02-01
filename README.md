## Run benchmarks

#### Generate plots:

> Draws plots for grpc and http benchmarks

usage: `go run  cmd/benchmark/plots/plots.go <number of requests> <number of concurrent request>`

```bash
go run cmd/benchmark/plots/plots.go 1000 20
```

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

### Notes:

- There are some benchmark tests in `request_benchmark_test.go` file. You can run them with `go test -bench=. -benchtime=1000x -benchmem`. There are 2 tests as `BenchmarkServer` and `BenchmarkServer_NewClientEachRequest`. I tested that what if it creates new gRPC connection for each request or use a single one for all requests. According to the results of those benchmarks, creating new connection for each request is not efficient for gRPC while it loses advantages of itself.

package main

import (
	"fmt"
	"os"
	"strconv"

	b "github.com/anilsenay/go-http-vs-grpc/cmd/benchmark/grpc/benchmark"
)

func main() {
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

	requestDurations := b.RunBenchmark(requestNumber, concurrency)
	fmt.Println(requestDurations)
}

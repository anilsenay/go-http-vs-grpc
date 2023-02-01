package grpc

import (
	"context"
	"fmt"

	"github.com/anilsenay/go-http-vs-grpc/grpc/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	RequestUrl string
	Client     *api.CategoryServiceClient
}

func NewGrpcClient(requestUrl string) *GrpcClient {
	conn, err := grpc.Dial(requestUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := api.NewCategoryServiceClient(conn)

	return &GrpcClient{
		RequestUrl: requestUrl,
		Client:     &client,
	}
}

func (g *GrpcClient) GetCategoryTree() {
	// start := time.Now()

	_, err := (*g.Client).GetCategoryTree(context.Background(), &api.CategoryTreeRequest{})
	if err != nil {
		panic(err)
	}

	// fmt.Println(len(resp.Categories))

	// if resp != nil {
	// 	recursivePrint(resp.Categories)
	// }

	// fmt.Printf("took: %d ns\n", time.Since(start).Nanoseconds())
}

func recursivePrint(r []*api.Category) {
	for _, c := range r {
		fmt.Println(c.Name)
		recursivePrint(c.SubCategories)
	}
}

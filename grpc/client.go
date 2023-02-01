package grpc

import (
	"context"

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

func (g *GrpcClient) GetCategoryTree() (*api.CategoryTreeResponse, error) {
	resp, err := (*g.Client).GetCategoryTree(context.Background(), &api.CategoryTreeRequest{})
	if err != nil {
		return nil, err
	}

	return resp, err
}

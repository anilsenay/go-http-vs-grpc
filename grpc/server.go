package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/anilsenay/go-http-vs-grpc/grpc/api"

	"google.golang.org/grpc"
)

// to code generation:
// cd grpc
// protoc -I=. --go_out=api --go_opt=paths=source_relative --go-grpc_out=api --go-grpc_opt=paths=source_relative category.proto

type GrpcServer struct {
	Host       string
	Port       string
	GrpcServer *grpc.Server
	Listener   net.Listener
}

func NewGrpcServer(host, port string) *GrpcServer {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		panic(err)
	}

	return &GrpcServer{
		Listener:   lis,
		GrpcServer: grpc.NewServer(),
		Host:       host,
		Port:       port,
	}
}

type categoryService struct {
	api.UnimplementedCategoryServiceServer
}

func (s *categoryService) GetCategoryTree(ctx context.Context, in *api.CategoryTreeRequest) (*api.CategoryTreeResponse, error) {
	// log.Printf("Received: %v", in.GetRoot())
	return &ExampleResponse, nil
}

func (g *GrpcServer) Serve() {
	s := grpc.NewServer()
	api.RegisterCategoryServiceServer(s, &categoryService{})
	log.Printf("server listening at %v", g.Listener.Addr())
	if err := s.Serve(g.Listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	g.Listener.Close()
}

func (g *GrpcServer) Shutdown() {
	g.GrpcServer.GracefulStop()
}

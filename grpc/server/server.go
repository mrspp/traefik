package main

import (
	"context"
	"errors"
	"fmt"
	hello "grpc-load-balancing/services/proto"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type server struct {
	address string
}

func (s *server) SayHello(ctx context.Context, request *hello.HelloReq) (*hello.HelloRes, error) {
	name := request.Name
	response := &hello.HelloRes{
		Mes: "Im" + name,
	}
	return response, nil
}

func (s *server) HelloBack(ctx context.Context, req *hello.HelloReq) (*hello.HelloBackRes, error) {
	name := req.Name
	response := &hello.HelloBackRes{
		Mes: "Hi" + name,
		Adj: "From" + s.address,
	}
	return response, nil
}

func main() {

	address := os.Getenv("ADDRESS")

	if address == "" {
		panic(errors.New("ADDRESS enviroment variable is required"))
	}

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: 10 * time.Second,
		}),
	}

	service := &server{address: address}

	grpcServer := grpc.NewServer(opts...)

	hello.RegisterHelloServiceServer(grpcServer, service)

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v", address)

	// s := grpc.NewServer()
	// hello.RegisterHelloServiceServer(s, &server{})
	grpcServer.Serve(lis)
}

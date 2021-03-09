package main

import (
	"context"
	"fmt"
	hello "grpc-load-balancing/services/proto"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct{}

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
		Adj: "Are you ok?",
	}
	return response, nil
}

func main() {

	address := os.Getenv("ADDRESS")
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v", address)

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &server{})
	s.Serve(lis)
}

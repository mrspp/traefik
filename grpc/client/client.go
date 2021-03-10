package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	hello "grpc-load-balancing/services/proto"
)

func main() {
	conn, err := grpc.Dial("dns://localhost:1053/my-grpc-loadbalance.demo.local:5000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithTimeout(10*time.Second),
	)
	if err != nil {
		panic(err)
	}

	client := hello.NewHelloServiceClient(conn)

	for {
		res, err := client.HelloBack(context.Background(), &hello.HelloReq{Name: " Josh"})
		if err != nil {
			panic(err)
		}
		fmt.Println("server address:", res.Mes, res.Adj)
		time.Sleep(time.Second)
	}
}

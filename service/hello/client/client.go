package main

import (
	"context"
	"fmt"
	"log"
	hello "traefik/service/hello/proto"

	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:9000", opts)

	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()
	client := hello.NewHelloServiceClient(cc)
	request := &hello.HelloReq{Name: "Ken"}
	resp, _ := client.SayHello(context.Background(), request)
	resp2, _ := client.HelloBack(context.Background(), request)

	fmt.Printf("Receive response => [%v]\n", resp.Mes)
	fmt.Printf("Receive response => [%v. %v]\n", resp2.Mes, resp2.Adj)

}

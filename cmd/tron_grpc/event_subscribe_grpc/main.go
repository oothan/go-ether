package main

import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	server := client.NewGrpcClientWithTimeout("grpc.trongrid.io:50051", 5*time.Second)
	server.SetAPIKey("464c4472-6063-4a8b-8b63-4e17c2363878")

	if err := server.Start(grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		panic(err)
	}

	fmt.Println("Success connected")
	//
	//subscriber, _ := goczmq.NewDealer("grpc.trongrid.io:50051")
	////server.Conn.Close()
	//
	//subscriber.SetSubscribe("contractLogTrigger")
	//
	//for {
	//	msg, _ := subscriber.RecvMessage()
	//	fmt.Println(msg)
	//}

}

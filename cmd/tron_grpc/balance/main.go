package main

import (
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	server := client.NewGrpcClientWithTimeout("grpc.trongrid.io:50051", 5*time.Second)
	server.SetAPIKey("464c4472-6063-4a8b-8b63-4e17c2363878")

	if err := server.Start(grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		panic(err)
	}

	fmt.Println("Success connected")

	acc, err := server.GetAccount("TNzvgxSEJsp5HsbgXSGrXUGvAd1p487J9d")
	if err != nil {
		panic(err)
	}

	fmt.Println("Balance :", acc.GetBalance())

	bal, err := server.TRC20ContractBalance("TRyzLGxGAhNyBNckkotEALznNd7sJVezeQ", "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")
	if err != nil {
		panic(err)
	}
	fmt.Println("USDT Balance :", bal)

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(6))))
	b, _ := value.Float64()
	fmt.Println("balance ", b)

	//server.TRC20Send("", "", "", )

	server.Conn.Close()

	//github issue
	// https://github.com/tronprotocol/java-tron/issues/2673
	// https://github.com/tronprotocol/TIPs/blob/master/tip-28.md
}

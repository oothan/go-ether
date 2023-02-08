package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-ether-dev/logger"
	"log"
)

func main() {
	baseUrl := fmt.Sprintf("https://mainnet.infura.io/v3/%s", "aff096ef309945c2841a5bd4c555a3c6")
	client, err := ethclient.Dial(baseUrl)
	if err != nil {
		logger.Sugar.Error(err)
		log.Fatalf("failed to connect to ether RPC server: %v:\n", err)
	}

	hash := common.HexToHash("0xedcc85dcc811625f9c3fe4f596b298f8d53b5b6e54b4fb8beee0e921472e8ca8")
	resp, isPending, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		logger.Sugar.Error(err)
	}

	logger.Sugar.Debugf("%+v", resp)
	logger.Sugar.Debug(isPending)

	rec, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	logger.Sugar.Debugf("\n%+v", rec)
}

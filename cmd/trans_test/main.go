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

	hash := common.HexToHash("0x6f17ff58cdc246ca48d8c17c16d1398b251acd3185b3c252e28817a55129fdc3")
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

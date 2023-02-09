package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	token "go-ether-dev/contracts_erc20"
	"go-ether-dev/logger"
	"log"
	"math"
	"math/big"
)

func main() {
	baseUrl := fmt.Sprintf("https://mainnet.infura.io/v3/%s", "aff096ef309945c2841a5bd4c555a3c6")
	client, err := ethclient.Dial(baseUrl)
	if err != nil {
		logger.Sugar.Error(err)
		log.Fatalf("failed to connect to ether RPC server: %v:\n", err)
	}

	////hash := common.HexToHash("0xedcc85dcc811625f9c3fe4f596b298f8d53b5b6e54b4fb8beee0e921472e8ca8")
	//hash := common.HexToHash("0x49df63f80f476aa2ed31bbc59b9da92ad5e48ddca9e3a2d8099ec5bd9424480d")
	//resp, isPending, err := client.TransactionByHash(context.Background(), hash)
	//if err != nil {
	//	logger.Sugar.Error(err)
	//}
	//
	//logger.Sugar.Debugf("%+v", resp)
	//logger.Sugar.Debug(isPending)
	//
	//rec, err := client.TransactionReceipt(context.Background(), hash)
	//if err != nil {
	//	logger.Sugar.Error(err)
	//	return
	//}
	//
	//logger.Sugar.Debugf("\n%+v", rec)

	tokenAddress := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		return
	}

	userAddr := common.HexToAddress("osmo19l9ym50dl8dnmu8um6sd9t0cps8027rqvev5vhew")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, userAddr)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	fbal := new(big.Float)
	fbal.SetString(bal.String())

	balance := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	b, _ := balance.Float64()

	logger.Sugar.Debug(b)
}

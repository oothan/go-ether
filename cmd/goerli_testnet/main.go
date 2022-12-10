package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	token "go-ether-dev/contracts_erc20"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type logTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type logApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func main() {
	client, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/164ee0f7f9fa4f94905c5f11e90ec1e9")
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("error on subscribing : %v\n", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	for {
		select {
		case err := <-sub.Err():
			fmt.Printf(err.Error())

		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				var transferEvt logTransfer

				err := contractAbi.UnpackIntoInterface(&transferEvt, "Transfer", vLog.Data)
				if err != nil {
					fmt.Println(err)
				}

				transferEvt.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvt.To = common.HexToAddress(vLog.Topics[2].Hex())
				customerAddress := transferEvt.From.Hex()
				bankAddress := transferEvt.To.Hex()
				hash := vLog.TxHash.Hex()

				fbal := new(big.Float)
				fbal.SetString(transferEvt.Tokens.String())
				value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(6)))
				//fmt.Println("balance: ", value) // "balance: 74605500.647409"

				i, _ := value.Float64()

				if strings.ToUpper("0xc3ec58d1810f96a26ff7ea2aef2df61fec3396f9") == strings.ToUpper(customerAddress) ||
					strings.ToUpper("0x3f98d756c755b77212F2A03E852a99B633a5061e") == strings.ToUpper(bankAddress) {
					fmt.Println(customerAddress)
					fmt.Println(bankAddress)
					fmt.Println(hash)
					fmt.Println(i)
				}

				if strings.ToUpper("0xc3ec58d1810f96a26ff7ea2aef2df61fec3396f9") == strings.ToUpper(bankAddress) ||
					strings.ToUpper("0x3f98d756c755b77212F2A03E852a99B633a5061e") == strings.ToUpper(customerAddress) {
					fmt.Println(customerAddress)
					fmt.Println(bankAddress)
					fmt.Println(hash)
				}

				//fmt.Println(customerAddress)
				//fmt.Println(bankAddress)
				//fmt.Println(hash)
				//fmt.Println(i)

			}
		}
	}
}

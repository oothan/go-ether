package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-ether-dev/token"
	"log"
	"math/big"
	"strings"
)

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func main() {
	//client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/2e36f270c48945bfade2661724b52075")
	//if err != nil {
	//	log.Error(err.Error())
	//}

	//contractAddress := common.HexToAddress("0xF7931B9b1fFF5Fc63c45577C43DFc0D0dEf16C46")
	//query := ethereum.FilterQuery{
	//	FromBlock: big.NewInt(2394201),
	//	ToBlock:   big.NewInt(2394201),
	//	Addresses: []common.Address{
	//		contractAddress,
	//	},
	//}
	//
	//logs, err := client.FilterLogs(context.Background(), query)
	//if err != nil {
	//	log.Error(err.Error())
	//}
	//
	//contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	//if err != nil {
	//	log.Error(err.Error())
	//}
	//
	//for _, vLog := range logs {
	//	fmt.Println("1==>", vLog.BlockHash.Hex())
	//	fmt.Println("2==>", vLog.BlockNumber)
	//	fmt.Println("3==>", vLog.TxHash.Hex())
	//
	//	//event := struct {
	//	//	Key   [32]byte
	//	//	Value [32]byte
	//	//}{}
	//	res, err := contractAbi.Unpack("ItemSet", vLog.Data)
	//	if err != nil {
	//		log.Error(err.Error())
	//	}
	//
	//	fmt.Println("4==>", res)
	//
	//	var topics [4]string
	//	for i := range vLog.Topics {
	//		topics[i] = vLog.Topics[i].Hex()
	//	}
	//	fmt.Println("5==>", topics[0])
	//	fmt.Println("6==>", topics[1])
	//	fmt.Println("7==>", topics[2])
	//	fmt.Println("8==>", topics[3])
	//}
	//
	//eventSignature := []byte("ItemSet(bytes32,bytes32)")
	//hash := crypto.Keccak256Hash(eventSignature)
	//fmt.Println("9==>", hash.Hex())

	//address := common.HexToAddress("0x2c8ecdca169a3b553b766d341eba6636d792e595cf8202186f5d6f0e0a8eb486")
	//instance, err := store.NewStore(address, client)
	//if err != nil {
	//	log.Error(err.Error())
	//}
	//
	//fmt.Println("contract is loaded")
	//_ = instance
	//
	//fmt.Println()

	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal("1 ", err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")
	address := common.HexToAddress("0xF7931B9b1fFF5Fc63c45577C43DFc0D0dEf16C46")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(15967117),
		ToBlock:   big.NewInt(15967117),
		Addresses: []common.Address{
			address,
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal("2 ", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal("3 ", err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("111 ==> Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("222 ==> Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("333 ==> Log Name: Transfer\n")

			var transferEvent LogTransfer

			_, err := contractAbi.Unpack("Transfer", vLog.Data)
			if err != nil {
				log.Fatal("4 ", err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

		case logApprovalSigHash.Hex():
			fmt.Printf("444 ==> Log Name: Approval\n")

			var approvalEvent LogApproval

			_, err := contractAbi.Unpack("Approval", vLog.Data)
			if err != nil {
				log.Fatal("5 ", err)
			}

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
		}

		fmt.Printf("\n\n")
	}

}

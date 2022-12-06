package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"math/big"
)

func main() {
	hash, _ := ETHTransferFrom(
		8,
		"",
		"",
		"",
	)
	fmt.Println(hash)
}

func ETHTransferFrom(balance float64, _fromAddress, _toAddress, _privateKey string) (string, error) {
	fmt.Println(" ............... Starting Transferring .............. ")
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println(ok)
		return "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("Address B ", fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	fromAddress = common.HexToAddress(_fromAddress)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Gas Price ", gasPrice)

	toAddress := common.HexToAddress(_toAddress)
	tokenAddress := common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7") // @todo 授权钱包地址,需要改成参数

	transferFnSignature := []byte("transferFrom(address,address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	paddedAddressFrom := common.LeftPadBytes(fromAddress.Bytes(), 32)

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	bal := fmt.Sprintf("%v000000", int64(balance))
	amount := new(big.Int)
	amount.SetString(bal, 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	//data = append(data, []byte("0x")...)
	data = append(data, methodID...)
	data = append(data, paddedAddressFrom...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     fromAddress,
		To:       &toAddress,
		Data:     data,
		Value:    value,
		GasPrice: gasPrice,
	})
	if err != nil {
		return "", err
	}
	fmt.Println("Gas Limit ", gasLimit)

	// gasLimit = gasLimit + gasLimit + gasLimit
	// fmt.Println()(gasLimit) // 23256
	//gasLimit = 1300000 //1.9 usdt
	gasLimit = 2100000 //2.5 usdt

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	fmt.Println(tx)
	//chainID, err := client.NetworkID(context.Background())
	chainID := big.NewInt(1) // chain id ==> prod
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("tx sent: ", signedTx.Hash().Hex())

	return signedTx.Hash().Hex(), nil
}

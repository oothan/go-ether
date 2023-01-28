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
	"log"
	"math/big"
)

func main() {
	hash, _ := tt(
		11,
		"0x92e1f9dbD7D95a2eC3bFb1665744cF46b1f34C81",
		"0xe9919fe1e8fdfcd504b9526d88749e96f8cb3acf",
		"bbda3906694dce3ef1f93f9a6e775159c22a333edf03b25b28bfe3174fc6fbe6",
	)
	fmt.Println(hash)
}

func tt(balance float64, _fromAddress, _toAddress, _privateKey string) (string, error) {
	ERC20USDTContractAddr := "0xdac17f958d2ee523a2206206994597c13d831ec7"
	baseUrl := fmt.Sprintf("https://mainnet.infura.io/v3/%s", "2e36f270c48945bfade2661724b52075")
	client, err := ethclient.Dial(baseUrl)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("failed to connect to ether RPC server: %v:\n", err)
	}

	ctx := context.Background()
	privateKey, err := crypto.HexToECDSA(_privateKey) // @todo 授权钱包私钥,需要改成参数
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println(err)
		return "", nil
	}

	addr := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, addr)
	if err != nil {
		return "", err
	}

	value := big.NewInt(0) // in wei (0 eth)

	//convert to common.Address
	fromAddress := common.HexToAddress(_fromAddress)
	toAddress := common.HexToAddress(_toAddress)
	contractAddr := common.HexToAddress(ERC20USDTContractAddr)

	transferFnSignature := []byte("transferFrom(address,address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddressFrom := common.LeftPadBytes(fromAddress.Bytes(), 32)
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	//logger.Sugar.Debug(hexutil.Encode(paddedAddress))

	bal := fmt.Sprintf("%v000000", int64(balance))
	amount := new(big.Int)
	amount.SetString(bal, 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddressFrom...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Gas Price :", gasPrice)

	/*
		gasPrice := big.NewInt(100000)
		var i, e = big.NewInt(10), big.NewInt(9) // 9 decimal
		i.Exp(i, e, nil)
		gasPrice.Mul(gasPrice, i)
		logger.Sugar.Debug("gas price : ", gasPrice)
	*/

	gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:     fromAddress,
		To:       &toAddress,
		Data:     data,
		Value:    value,
		GasPrice: gasPrice,
	})
	if err != nil {
		return "", err
	}
	gasLimit = gasLimit * 5
	fmt.Println("Gas Limit :", gasLimit)

	tx := types.NewTransaction(nonce, contractAddr, value, gasLimit, gasPrice, data)

	// logger.Sugar.Debug(tx)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("tx sent: ", signedTx.Hash().Hex())

	return signedTx.Hash().Hex(), nil
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
	//gasPrice.Mul(gasPrice, big.NewInt(2))
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
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Gas Limit ", gasLimit)

	//gasLimit = gasLimit + gasLimit + gasLimit
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

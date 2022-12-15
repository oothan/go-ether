package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
	"log"
	"math"
	"math/big"
	"strings"
)

// var baseUrl = "https://cloudflare-eth.com"
var baseUrl = "https://goerli.infura.io/v3/164ee0f7f9fa4f94905c5f11e90ec1e9"
var wsUrl = "wss://goerli.infura.io/ws/v3/164ee0f7f9fa4f94905c5f11e90ec1e9"
var usdtToken = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
var myAddr = "0x64d17712Fc5795e0784bbA04f8dB83Fe16E23f17"

func main() {
	client, err := ethclient.Dial(baseUrl)
	if err != nil {
		panic(err)
	}

	conn, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//val := map[string]interface{}{
	//	"jsonrpc": "2.0",
	//	"id":      1,
	//	"method":  "eth_subscribe",
	//	"params": []interface{}{
	//		"logs",
	//		map[string]interface{}{
	//			"fromBlock": "latest",
	//			"toBlock":   "latest",
	//		},
	//	},
	//}

	val := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_subscribe",
		"params": []interface{}{
			"newPendingTransactions",
		},
	}
	if err := conn.WriteJSON(val); err != nil {
		panic(err)
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		//fmt.Println(string(msg))
		//fmt.Println()

		//resp := &PendindResp{}
		//resp.Params = &Result{}
		//resp.Params.Result = &ParamResp{}
		//if err := json.Unmarshal(msg, &resp); err != nil {
		//	panic(err)
		//}
		//
		////fmt.Println(resp.Params)
		//
		////fmt.Println(resp)
		//
		//hash := common.HexToHash(resp.Params.Result.TransactionHash)
		//tx, isPending, err := client.TransactionByHash(context.Background(), hash)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//
		////fmt.Println(hash)
		////fmt.Println(tx)
		////fmt.Println(isPending)
		//
		//if tx != nil {
		//	chainID, err := client.NetworkID(context.Background())
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	i := big.NewInt(1)
		//
		//	if msg1, err := tx.AsMessage(types.NewEIP2930Signer(chainID), i); err != nil {
		//		fmt.Println("From Address : ", msg1.From().Hex())
		//		fmt.Println("To Address : ", msg1.To().Hex())
		//		fmt.Println("Value : ", msg1.Value())
		//
		//		//if strings.ToUpper("0xc3ec58d1810f96a26ff7ea2aef2df61fec3396f9") == strings.ToUpper(msg1.From().Hex()) ||
		//		//	strings.ToUpper("0x3f98d756c755b77212F2A03E852a99B633a5061e") == strings.ToUpper(msg1.To().Hex()) {
		//		//	fmt.Println("<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		//		//	fmt.Println("From Address : ", msg1.From().Hex())
		//		//	fmt.Println("To Address : ", msg1.To().Hex())
		//		//	fmt.Println("Value : ", msg1.Value())
		//		//
		//		//	fbal := new(big.Float)
		//		//	fbal.SetString(msg1.Value().String())
		//		//	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(18))))
		//		//	fmt.Println("Value : ", value)
		//		//}
		//		//
		//		//if strings.ToUpper("0xc3ec58d1810f96a26ff7ea2aef2df61fec3396f9") == strings.ToUpper(msg1.To().Hex()) ||
		//		//	strings.ToUpper("0x3f98d756c755b77212F2A03E852a99B633a5061e") == strings.ToUpper(msg1.From().Hex()) {
		//		//	fmt.Println("<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		//		//	fmt.Println("From Address : ", msg1.From().Hex())
		//		//	fmt.Println("To Address : ", msg1.To().Hex())
		//		//	fmt.Println("Value : ", msg1.Value())
		//		//
		//		//	fbal := new(big.Float)
		//		//	fbal.SetString(msg1.Value().String())
		//		//	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(18))))
		//		//	fmt.Println("Value : ", value)
		//		//}
		//
		//		fmt.Println("isPending :", isPending, ", hash :", hash.Hex())
		//		fmt.Println()
		//	}
		//}

		resp := &PendindResp{}
		resp.Params = &ParamResp{}
		if err := json.Unmarshal(msg, &resp); err != nil {
			panic(err)
		}

		hash := common.HexToHash(resp.Params.Result)
		tx, isPending, err := client.TransactionByHash(context.Background(), hash)
		if err != nil {
			fmt.Println(err)
		}

		if tx != nil {
			chainID, err := client.NetworkID(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			i := big.NewInt(1)

			if msg1, err := tx.AsMessage(types.NewEIP2930Signer(chainID), i); err != nil {
				//fmt.Println("From Address : ", msg1.From().Hex())
				//fmt.Println("To Address : ", msg1.To().Hex())
				//fmt.Println("Value : ", msg1.Value())

				if msg1.To() != nil {
					if strings.ToUpper("0x92e1f9dbD7D95a2eC3bFb1665744cF46b1f34C81") == strings.ToUpper(msg1.To().Hex()) ||
						strings.ToUpper("0x3f98d756c755b77212f2a03e852a99b633a5061e") == strings.ToUpper(msg1.To().Hex()) {
						fmt.Println("<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
						fmt.Println("From Address : ", msg1.From().Hex())
						fmt.Println("To Address : ", msg1.To().Hex())
						fmt.Println("Value : ", msg1.Value())

						fbal := new(big.Float)
						fbal.SetString(msg1.Value().String())
						value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(18))))
						fmt.Println("Value : ", value)
						fmt.Println("isPending :", isPending, ", hash :", hash.Hex())
					}
				}

			}
			//
			fmt.Println()
		}

	}
}

type PendindResp struct {
	Params *ParamResp `json:"params"`
}

type ParamResp struct {
	Result string `json:"result"`
}

//type PendindResp struct {
//	Params  *Result `json:"params"`
//	Jsonrpc string  `json:"jsonrpc"`
//	Method  string  `json:"method"`
//}
//
//type Result struct {
//	Result       *ParamResp `json:"result"`
//	Subscription string     `json:"subscription"`
//}
//
//type ParamResp struct {
//	TransactionHash string `json:"transactionHash"`
//}

/*
func main() {
  val := map[string]interface{}{
    "jsonrpc": "2.0",
    "method":  "eth_getBalance",
    "params": []string{
      "0xF7931B9b1fFF5Fc63c45577C43DFc0D0dEf16C46",
      "latest",
    },
    "id": 1,
  }

  payloadBuf := new(bytes.Buffer)
  json.NewEncoder(payloadBuf).Encode(val)

  req, err := http.NewRequest(http.MethodPost, baseUrl, payloadBuf)
  if err != nil {
    panic(err)
  }
  req.Header.Add("Content-Type", "application/json; charset=UTF-8")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  data := &RespData{}
  if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
    panic(err)
  }

  fbalance := new(big.Float)
  fbalance.SetString(data.Result)
  ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

  fmt.Println(ethValue)
}

type RespData struct {
  Result string json:"result"
}
*/

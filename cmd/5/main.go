package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TronApprovalList struct {
	Total uint64     `json:"total"`
	Data  []Approval `json:"data"`
}

type Approval struct {
	Amount          string    `json:"amount"`
	Unlimited       bool      `json:"unlimited"`
	ToAddress       string    `json:"to_address"`
	ContractAddress string    ` json:"contract_address"`
	FromAddress     string    `json:"from_address"`
	TokenInfo       TokenInfo `json:"tokenInfo"`
}

type TokenInfo struct {
	TokenId      string `json:"tokenId"`
	TokenAbbr    string `json:"tokenAbbr"`
	TokenName    string `json:"tokenName"`
	TokenDecimal int64  `json:"tokenDecimal"`
}

func main() {
	url := "https://apilist.tronscanapi.com/api/account/approve/list?address=TRnuSWKUtuwKCGbrFZ6UYsDLwujjnZ2t4m&limit=50&start=0&type=project"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	approvalList := TronApprovalList{}
	err = json.NewDecoder(resp.Body).Decode(&approvalList)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%+v\n", approvalList)

}

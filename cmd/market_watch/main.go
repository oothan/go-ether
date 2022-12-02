package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/gorilla/websocket"
)

// Initialize a connection using your API key
// You can generate an API key here: https://cryptowat.ch/account/api-access
// Paste your API key here:
const (
	APIKEY = "TDJMN1GUPCN1UIOFTADM"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("wss://stream.cryptowat.ch/connect?apikey="+APIKEY, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Read first message, which should be an authentication response
	_, message, err := c.ReadMessage()
	var authResult struct {
		AuthenticationResult struct {
			Status string `json:"status"`
		} `json:"authenticationResult"`
	}
	err = json.Unmarshal(message, &authResult)
	if err != nil {
		panic(err)
	}
	// Send a JSON payload to subscribe to a list of resources
	// Read more about resources here: https://docs.cryptowat.ch/websocket-api/data-subscriptions#resources
	resources := []string{
		"instruments:231:summary", // btcusdt
		"instruments:165:summary", // ethusdt
		"instruments:152:summary", // usdtusd
		"instruments:636:summary", // trxusdt

	}
	subMessage := struct {
		Subscribe SubscribeRequest `json:"subscribe"`
	}{}

	// No map function in golang :-(
	for _, resource := range resources {
		subMessage.Subscribe.Subscriptions = append(
			subMessage.Subscribe.Subscriptions,
			Subscription{
				StreamSubscription: StreamSubscription{
					Resource: resource,
				},
			},
		)
	}
	msg, err := json.Marshal(subMessage)
	err = c.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		panic(err)
	}

	// Process incoming BTC/USD trades
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal("Error reading from connection", err)
			return
		}

		//var update Update
		ticker := &Ticker{}
		err = json.Unmarshal(message, &ticker)
		if err != nil {
			panic(err)
		}
		//fmt.Println(string(message))

		//fmt.Println(ticker)

		fmt.Println()

		percent, _ := strconv.ParseFloat(ticker.MarketUpdate.SummaryUpdate.ChangePercentStr , 64)
		percent = RoundFloat(percent, 4)
		if ticker.MarketUpdate.Market.CurrencyPairId == "231" {
			log.Printf(

				"BTC/USDT | LastPrice: %v | ChangePercentage: %v%v \n",
				ticker.MarketUpdate.SummaryUpdate.LastStr,
				percent * 100,
				"%",
				)
		} else if ticker.MarketUpdate.Market.CurrencyPairId == "165" {
			log.Printf(
				"ETH/USDT | LastPrice: %v | ChangePercentage: %v%v \n",
				ticker.MarketUpdate.SummaryUpdate.LastStr,
				percent,
				"%",
			)
		} else if ticker.MarketUpdate.Market.CurrencyPairId == "152" {
			log.Printf(
				"USDT/USD | LastPrice: %v | ChangePercentage: %v%v \n",
				ticker.MarketUpdate.SummaryUpdate.LastStr,
				percent,
				"%",
			)
		} else if ticker.MarketUpdate.Market.CurrencyPairId == "636" {
			log.Printf(
				"TRX/USDT | LastPrice: %v | ChangePercentage: %v%v \n",
				ticker.MarketUpdate.SummaryUpdate.LastStr,
				percent,
				"%",
			)
		}
	}
}

// Helper types for JSON serialization

type Subscription struct {
	StreamSubscription `json:"streamSubscription"`
}

type StreamSubscription struct {
	Resource string `json:"resource"`
}

type SubscribeRequest struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

type Update struct {
	MarketUpdate struct {
		Market struct {
			MarketId int `json:"marketId,string"`
			ExternalId     string `json:"externalId"`
		} `json:"market"`

		//TradesUpdate struct {
		//	Trades []Trade `json:"trades"`
		//} `json:"tradesUpdate"`
	} `json:"marketUpdate"`

	SummaryUpdate struct{
		LastStr string `json:"lastStr"`
		ChangePercentStr string `json:"changePercentStr"`
	} `json:"summaryUpdate"`
}

type Summary struct {
	LastStr string `json:"lastStr"`
	ChangePercentStr string `json:"changePercentStr"`
}

type Trade struct {
	Timestamp     int `json:"timestamp,string"`
	TimestampNano int `json:"timestampNano,string"`

	ExternalId     string `json:"externalId"`
	AmountQuoteStr string `json:"amountQuoteStr"`
	Price          string `json:"priceStr"`
	Amount         string `json:"amountStr"`
}

type Ticker struct {
	MarketUpdate MarketUpdate `json:"marketUpdate"`

}

type MarketUpdate struct {
 	Market Market `json:"market"`
	SummaryUpdate SummaryUpdate `json:"summaryUpdate"`
}

type Market struct {
	CurrencyPairId string `json:"currencyPairId"`
}

type SummaryUpdate struct {
	LastStr string `json:"lastStr"`
	ChangePercentStr string `json:"changePercentStr"`
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}


package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// for demo

func main() {

	c := colly.NewCollector()
	c.UserAgent = "Go program"

	c.OnRequest(func(r *colly.Request) {

		for key, value := range *r.Headers {
			fmt.Printf("111 ==> %s: %s\n", key, value)
		}

		fmt.Println("222 ==> ", r.Method)
	})

	var count int64
	c.OnHTML("tr td a", func(e *colly.HTMLElement) {

		// fmt.Println(e.DOM.Nodes[0], "DD", e.DOM.Nodes[0].FirstChild.Data == " ")
		if e.DOM.HasClass("hash-tag text-truncate") && e.DOM.Nodes[0].FirstChild.Type != 3 && e.Text != "USDT" {
			fmt.Println(e)
			fmt.Println("-----------------------------")
			fmt.Println("333 ==> ", e.Text, "DDDD")
			count++
		}

	})

	c.OnResponse(func(r *colly.Response) {

		fmt.Println("-----------------------------")

		fmt.Println(r.StatusCode)

		for key, value := range *r.Headers {
			fmt.Printf("%s: %s\n", key, value)
		}
	})

	c.Visit("https://etherscan.io/tokenapprovalchecker?search=0xF7931B9b1fFF5Fc63c45577C43DFc0D0dEf16C46")
	fmt.Println("555 ==> ", count)
}

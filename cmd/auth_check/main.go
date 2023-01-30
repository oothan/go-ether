package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

func main() {
	c := colly.NewCollector()
	c.UserAgent = "Go program"
	c.SetRequestTimeout(time.Second * 120)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 10 * time.Second,
	})

	var count int64
	existFlag := false
	c.OnHTML("table#mytable > tbody", func(e *colly.HTMLElement) {

		// fmt.Println(e.DOM.Nodes[0], "DD", e.DOM.Nodes[0].FirstChild.Data == " ")
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			if el.ChildText("td:nth-child(3)") == "Tether USD" {
				count++
			}

			if strings.EqualFold(el.ChildText("td:nth-child(4)"), "0x0F188fff97ba7E19f988B6c847f26D48AB0509d7") {
				existFlag = true
				fmt.Println(el.ChildText("td:nth-child(4)"))
			}
		})

	})

	c.Visit(fmt.Sprintf("https://etherscan.io/tokenapprovalchecker?search=0xd846e03b2c847ca73e28eb7559a6302170e0e35f"))
	c.Wait()
	fmt.Println("555 ==> ", count)
	fmt.Println(existFlag)
}

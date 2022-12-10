package main

import (
	"fmt"
	"github.com/gocolly/colly"
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
	c.OnHTML("table#mytable > tbody", func(e *colly.HTMLElement) {

		// fmt.Println(e.DOM.Nodes[0], "DD", e.DOM.Nodes[0].FirstChild.Data == " ")
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			if el.ChildText("td:nth-child(3)") == "Tether USD" {
				count++
			}
		})

	})

	c.Visit(fmt.Sprintf("https://etherscan.io/tokenapprovalchecker?search=0x92e1f9dbd7d95a2ec3bfb1665744cf46b1f34c81"))
	c.Wait()
	fmt.Println("555 ==> ", count)
}

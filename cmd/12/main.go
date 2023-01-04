package main

import (
	"fmt"
	"time"
)

func main() {
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	loc, _ := time.LoadLocation(time.Now().Location().String())
	now := time.Now().In(loc)
	since := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, loc)
	until := time.Date(now.Year(), now.Month(), now.Day()-1, 23, 59, 59, 59, loc)
	//d := now.Sub(midnight)

	//fmt.Println(d)
	//fmt.Println(midnight)
	fmt.Println(now.Day())

	fmt.Println(since)
	fmt.Println(until)
	fmt.Println(time.Now().Location().String())
}

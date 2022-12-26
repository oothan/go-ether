package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//layoutISO := "2006-01-02"
	//tt := time.Now()
	//fmt.Println( tt.Format(layoutISO))
	////t, _ := time.Parse(layoutISO, )
	////createdAt := fmt.Sprintf("date(created_at) = date(\"%v\")", t.Format(layoutISO))

	tt := 0.012
	fmt.Println(100 * tt)

	i := fmt.Sprintf("%.2f", (tt * 100))
	f, _ := strconv.ParseFloat(i, 2)
	fmt.Println(f)

}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

package main

import (
	"fmt"
	"math"
)

func main() {
	bal := 0.00000046
	//fmt.Println(RoundFloat(float64(bal), 8))
	fmt.Println(bal)
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

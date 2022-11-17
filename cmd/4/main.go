package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	t := 0.0000007
	//fmt.Println(fmt.Sprintf("%.6f", RoundFloat(t, 6)))
	fmt.Println(RoundFloat(t, 6))
	//fmt.Println(RoundFloat1(t, 6))
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func RoundFloat1(x float64, prec int) float64 {
	frep := strconv.FormatFloat(x, 'g', prec, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}

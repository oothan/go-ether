package main

import (
	"fmt"
	"math"
)

func main() {
	var users []string
	for i := range [200]int{} {
		users = append(users, fmt.Sprintf("user:%v", i))
	}

	fmt.Println(users)
	userSlicing(users, 7)

}

func userSlicing(users []string, length int) {
	var (
		start = 0
		//different  = 0
		//realLength = 0
	)
	total := len(users) / length
	remainder := len(users) % length
	fmt.Println(remainder, total)
	for i := 1; i <= length; i++ {
		fmt.Println("Start : ", start, ", total : ", total)
		fmt.Println(users[start:total])

		start += total
		total += total

		//fmt.Println("Start : ", start, ", total : ", total)
		//if remainder != 0 && i == length {
		//	fmt.Println(users[start:remainder])
		//} else {
		//	fmt.Println(users[start:total])
		//}
		//start += total
		//total += total

		fmt.Println(i)
	}
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func foo(channel, quit chan string, i int) {

	channel <- fmt.Sprintf("goroutine %d started!", i)
	for {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		quit <- fmt.Sprintf("goRoutine %d completed!", i)
	}
}
func main() {

	channel := make(chan string)
	quit := make(chan string)

	for i := 0; i < 3; i++ {
		go foo(channel, quit, i)
	}

	for {
		select {
		case update := <-channel:
			fmt.Println(update)
		case quit := <-quit:
			fmt.Println(quit)
			return
		}
	}
}

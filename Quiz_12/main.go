package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	select {
	case ch <- 42:
		fmt.Println("Value sent to channel")

	default:
		fmt.Println("Channel not ready to receive value")
	}

	select {
	case val := <-ch:
		fmt.Println("Value receive from channel :", val)
	default:
		fmt.Println("Channel not ready to send value")
	}
}

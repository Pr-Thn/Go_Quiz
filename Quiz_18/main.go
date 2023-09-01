package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("A")
		}
	}()
	go func() {
		for {
			fmt.Println("B")
		}
	}()
	go func() {
		for {
			fmt.Println("C")
		}
	}()

	time.Sleep(time.Millisecond)
}

package main

import (
	"fmt"
)

var counter int

func incrementCounter() {
	counter++
}

func main() {
	for i := 0; i < 100000; i++ {
		go incrementCounter()
	}

	// time.Sleep(time.Second)

	fmt.Print("Final counter value :", counter)
}

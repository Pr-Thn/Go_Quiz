package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func incrementCounter() {
	mutex.Lock() // Use mutex to Protect counter value
	counter++
	mutex.Unlock()
}

func main() {
	for i := 0; i < 10000; i++ {
		go incrementCounter()
	}
	/* When a Goroutine calls incrementCounter, it first acquires the lock using mutex.Lock(),
	ensuring that only one Goroutine can enter the critical section at a time.*/

	/* The defer mutex.Unlock()
	statement releases the lock when the function exits, even if an error occurs or a panic is triggered.*/

	/* The main function demonstrates how to use the Counter type with Goroutines.
	It launches multiple Goroutines to increment the counter concurrently */

	// time.Sleep(time.Second)
	fmt.Println("Final counter value :", counter)
}

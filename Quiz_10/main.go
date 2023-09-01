package main

import "fmt"

func main() {
	fmt.Println("Start")

	greeting := make(chan string)
	// greeting <- "Hello"
	fmt.Println(greeting)
	
	fmt.Println("Finish")
}

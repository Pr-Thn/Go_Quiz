package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	a = a[:0]
	fmt.Println(a, len(a), cap(a))
}

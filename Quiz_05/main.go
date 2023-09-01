package main

import "fmt"

func main() {
	x := 1
	fmt.Println("x before : ", x)

	{
		x := 2
		fmt.Println("x inside : ", x)
	}
	// it is var it not const
	fmt.Println("x after : ", x)

}

package main

import "fmt"

func recoverAndPrint() {
	if r := recover(); r != nil {
		fmt.Println("Recover : ", r)
	}
}

/*
Recover is a built-in function that regains control of a panicking goroutine.
Recover is only useful inside deferred functions. During normal execution,
a call to recover will return nil and have no other effect. If the current goroutine is panicking,
a call to recover will capture the value given to panic and resume normal execution.
*/
func divide(x, y int) int {
	defer recoverAndPrint()
	// A defer statement pushes a function call onto a list.
	// The list of saved calls is executed after the surrounding function returns.
	// Defer is commonly used to simplify functions that perform various clean-up actions.
	if y == 0 {
		panic("division by zero")
	}

	return x / y
}

func main() {
	fmt.Println(divide(10, 0))
	// fmt.Println(divide(10, 2))
}

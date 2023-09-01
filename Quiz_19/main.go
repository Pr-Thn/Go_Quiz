package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumGoroutine())
	go func() {

	}()
	go func() {

	}()
	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())

}

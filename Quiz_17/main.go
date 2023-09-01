package main

import (
	"fmt"
	"time"
)

type MyStruct struct {
	value int
}

func (ms MyStruct) AddStruct(other MyStruct) {
	ms.value += other.value
}

func (ms *MyStruct) AddPointer(other *MyStruct) {
	ms.value += other.value
}

func main() {
	n := 10000000
	start := time.Now()

	// test using struct receive
	ms1 := MyStruct{value: 0}
	for i := 0; i < n; i++ {
		ms1.AddStruct(MyStruct{value: 1})
	}

	elapsed := time.Since(start)
	fmt.Printf("Using Struct receiver took %s\n", elapsed)

	// test using pointer receive
	start = time.Now() // reset the start time
	ms2 := &MyStruct{value: 0}
	for i := 0; i < n; i++ {
		ms2.AddPointer(&MyStruct{value: 1})
	}

	elapsed = time.Since(start)
	fmt.Printf("Using Pointer receiver took %s\n", elapsed)
}

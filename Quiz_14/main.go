package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"Address,omitempty"`
}

func main() {
	p := Person{Name: "John", Age: 30}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error : ", err)
	}
	fmt.Println(string(b))
	fmt.Println(b) // return Byte[] code
}

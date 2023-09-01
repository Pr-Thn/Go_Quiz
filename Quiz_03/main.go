package main

import (
	"fmt"
	"strconv"
)

func main() {
	decimalString := "1.2"

	decimal32Bit, _ := strconv.ParseFloat(decimalString, 3)
	fmt.Println(decimal32Bit)
	decimal64Bit, _ := strconv.ParseFloat(decimalString, 6)
	fmt.Println(decimal64Bit)

	fmt.Println(decimal32Bit == decimal64Bit)
}

package main

import "fmt"

func squares(numbers [5]int) {
	for index, number := range numbers {
		numbers[index] = number * number
		fmt.Println(numbers)
	}
}

func main() {
	numbers := [5]int{0, 1, 2, 3, 4}
	squares(numbers)
	fmt.Println(numbers)
}

// Output = {0,1,2,3,4} because it has no returned

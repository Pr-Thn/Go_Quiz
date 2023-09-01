package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for i := range nums {
		fmt.Println(i)
		nums = append(nums, i)
	}

	fmt.Println(nums)
}

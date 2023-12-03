package main

import "fmt"

type Item struct {
	Name  string
	Price float64
}

func main() {
	items := map[string]Item{
		"apple":  {"Apple", 1.0},
		"banana": {"Banana", 2.0},
	}
	apple := items["apple"]
	apple.Price = 1.5
	fmt.Println(items["apple"].Price)

	//  Cannot directly assign a value to a struct field stored within a map.
	//  Need to perform the assignment in two steps

	// FIXME : Assuming "apple" is a key in the map
	if _, ok := items["apple"]; ok {
		apple := items["apple"]
		apple.Price = 2.99 // This will not work
		// You cannot directly assign like items["apple"].Price
		// Instead, you need to update the value and then assign it back to the map
		appleItem := items["apple"]
		appleItem.Price = 2.99
		items["apple"] = appleItem
	}

}

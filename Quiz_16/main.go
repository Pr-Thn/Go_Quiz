package main

import "fmt"

func allUnique(strs []string) bool {
	strMap := make(map[string]bool)
	for _, str := range strs {
		if strMap[str] {
			return false //  string is already in the map, so it's not unique
		}
		strMap[str] = true // mark the string as seen
	}
	return true // all string are unique
}

func main() {
	// Call allUnique with some sample inputs
	input1 := []string{"hello", "world", "foo", "bar"}
	fmt.Printf("%v : %t\n", input1, allUnique(input1)) // [hello world foo bar] : true

	input2 := []string{"foo", "bar", "baz", "foo"}
	fmt.Printf("%v : %t\n", input2, allUnique(input2)) // [foo bar barz foo] : false

	input3 := []string{}
	fmt.Printf("%v : %t\n", input3, allUnique(input3)) // [] : true

}

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func IncrementAge(p *Person) {
	p.Age++
}

func main() {
	p1 := &Person{"Alice", 25}
	p2 := p1 // this line mean *p1 = *p2 {Name,Age}
	IncrementAge(p2)
	// Increment add value to p.Age in *p2 where *p1 = *p2
	// Not the same pointer but look at the same value

	fmt.Println(*p1)
	fmt.Println(*p2)
	// P1 and P2 are not in the same pointers but why does p1 and p2 are the same
	fmt.Println(&p1)
	fmt.Println(&p2)
}

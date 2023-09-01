package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Person struct {
	Name      string    `validate:"required"`
	Age       int       `validate:"gte=0,lte=130"`
	Addresses []Address `validate:"dive"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	State  string `validate:"required,len=2"`
	Zip    string `validate:"required,len=5"`
}

func main() {
	// Create a new validator instance
	validate := validator.New()

	// Create a person with an invalid address
	person := Person{
		Name: "John Smith",
		Age:  35,
		Addresses: []Address{
			{
				Street: "123 Main St",
				City:   "Anytown",
				State:  "XY",
				Zip:    "123456",
			},
		},
	}
	// Validate the person
	err := validate.Struct(person)
	if err != nil {
		// Print validation error
		fmt.Println(err)
		return
	}

	// Person is valid
	fmt.Println("Person is valid")
}

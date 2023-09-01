package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Person struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       int    `validate:"required"`
}

func main() {
	validate := validator.New()

	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       150,
	}

	err := validate.Struct(person)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// Get the name of the field that failed validation
			field := err.Field()

			// Get the value that failed validation
			value := err.Value()

			// Get the validation tag that was used
			tag := err.Tag()

			// Get the validation message with the filed name
			fmt.Printf("Error :field %s fialed validation with value %v using tag '%s'\n")
		}
	}
}

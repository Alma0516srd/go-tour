// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create two error variables, one called ErrInvalidValue and the other
// called ErrAmountTooLarge. Provide the static message for each variable.
// Then write a function called checkAmount that accepts a float64 type value
// and returns an error value. Check the value for zero and if it is, return
// the ErrInvalidValue. Check the value for greater than $1,000 and if it is,
// return the ErrAmountTooLarge. Write a main function to call the checkAmount
// function and check the return error value. Display a proper message to the screen.
package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidValue = errors.New("invalid value: amount cannot be zero")

	ErrAmountTooLarge = errors.New("amount too large: cannot exceed $1000")
)

func checkAmount(amount float64) error {

	if amount == 0 {
		return ErrInvalidValue
	}

	if amount > 1000 {
		return ErrAmountTooLarge
	}

	return nil
}

func main() {
	testAmounts := []float64{0, 500, 1500, 999.99}

	for _, amount := range testAmounts {
		fmt.Printf("Checking amount: $%.2f ... ", amount)

		err := checkAmount(amount)

		switch err {
		case ErrInvalidValue:
			fmt.Println("Error: invalid value provided.")

		case ErrAmountTooLarge:
			fmt.Println("Error: amount exceeds maximum limit.")

		case nil:
			fmt.Println("OK, amount is valid.")

		default:

			fmt.Printf("Unknown error: %v\n", err)
		}
	}

	fmt.Println("All amounts processed successfully!")
}

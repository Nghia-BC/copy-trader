package main

import (
	"fmt"
	"log"
)

/**
Write a program which can compute the factorial of the given numbers.
The results should be printed in a comma-seperated sequence on the single line.
8 --> 40320
**/

func main() {
	// variable
	var input int
	fmt.Print("Please enter the number: ")
	// input from the console
	// _ means that it will not be used somewhere
	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Print("Please enter the number!")
	}

	output, err := calF(input)

	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("Factorial is: %d", output)

}

func calF(input int) (uint64, error) {
	var output uint64 = 1

	if input < 0 {
		return 0, fmt.Errorf("factorial for negativ values is not allowed")
	}

	if input == 0 {
		return 1, nil
	}

	for i := input; i > 0; i-- {
		output *= uint64(i)
	}

	return output, nil
}

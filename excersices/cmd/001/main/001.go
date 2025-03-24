package main

/*
Write a program to find all such numbers which are divisible by 7 but are not multiple of 5 between 2000 and 3200
The number obtained should be printed a comma seperated sequence on a single line
*/

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// variable
	data := cal(2000, 3200) // 1 line
	fmt.Println(data)       // 1 line
}

func cal(low, high int) string {
	var numbers []string // 1
	// loop
	for i := low; i <= high; i++ { // 1
		// if else
		if i%7 == 0 && i%5 != 0 { // 1
			// slice
			// strconv.Itoa convert int to string
			numbers = append(numbers, strconv.Itoa(i)) // 1
		}
	}

	return strings.Join(numbers, ",") // 1
}

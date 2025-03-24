package main

import (
	"fmt"
	"log"
)

/**

With a given integral number n, write a program to generate a map that contains (i, i*i)
such that is an integral number between 1 and n (both included),
and then the program should print the map with representation of the value

Suppose the following input is supplied to the program: 8
Then, the output should be:
map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]

*/

func main() {
	var input int

	fmt.Println("Please type the number: ")
	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatal("Error happened: ", err)
	}

	output := task(input)
	fmt.Printf("%v", output)

}

// why map take ram if we have a big map?
func task(input int) map[int]int {
	// create a map with size of input

	var numbers = make(map[int]int, input)

	for i := 1; i <= input; i++ {
		numbers[i] = i * i
	}
	return numbers
}

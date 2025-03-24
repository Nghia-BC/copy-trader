package main

import "fmt"

type Person struct {
	name string
}

func changeName1(per *Person) {
	per.name = "Khang2"
}

func main() {
	person := Person{"Khang"}
	fmt.Print("name: ", person.name)

	changeName(person)

	fmt.Print("Name: ", person.name)

	changeName1(&person)

	fmt.Print("Name: ", person.name)
}

func changeName(person Person) {
	person.name = "Khang1"
}

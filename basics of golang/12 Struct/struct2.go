package main

import "fmt"

// Define a struct type
type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {

	person := Person{
		Name:   "jai",
		Age:    30,
		Gender: "Male",
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Gender:", person.Gender)
}

package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person

	pers1.name = "Dilip"
	pers1.age = 45
	pers1.job = "Engineering"
	pers1.salary = 6000

	pers2.name = "kumar"
	pers2.age = 20
	pers2.job = "Marketing"
	pers2.salary = 4500

	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)
}

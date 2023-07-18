package main

import "fmt"

func main() {
	num1 := 10
	num2 := 4

	add := num1 + num2
	fmt.Println("Addition:", add)

	sub := num1 - num2
	fmt.Println("Subtraction:", sub)

	mul := num1 * num2
	fmt.Println("Multiplication:", mul)

	div := num1 / num2
	fmt.Println("Division:", div)

	mod := num1 % num2
	fmt.Println("Remainder:", mod)

	num1++
	fmt.Println("Increment:", num1)

	num2--
	fmt.Println("Decrement:", num2)
}

package main

import "fmt"

func main() {
	num1 := 10
	num2 := 5
	position := 2

	fmt.Println("Bitwise AND:", num1&num2)
	fmt.Println("Bitwise OR:", num1|num2)

	fmt.Println("Bitwise XOR:", num1^num2)

	fmt.Println("Bitwise NOT:", ^num1)
	fmt.Println("Left shift:", num1<<position)
	fmt.Println("Right shift:", num1>>position)
}

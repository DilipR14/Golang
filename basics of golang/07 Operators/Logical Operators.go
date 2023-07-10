package main

import "fmt"

func main() {
	num1 := 10
	num2 := 5
	str1 := "Hello"
	str2 := "World"

	// Logical AND
	fmt.Println("Logical AND:")
	fmt.Println(num1 > num2 && str1 == str2)
	fmt.Println(num1 > num2 && str1 != str2)
	fmt.Println(num1 < num2 && str1 == str2)

	// Logical OR
	fmt.Println("Logical OR:")
	fmt.Println(num1 > num2 || str1 == str2)
	fmt.Println(num1 > num2 || str1 != str2)
	fmt.Println(num1 < num2 || str1 == str2)

	// Logical NOT
	fmt.Println("Logical NOT:")
	fmt.Println(!(num1 > num2))
	fmt.Println(!(num1 < num2))
}

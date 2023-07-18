package main

import "fmt"

func main() {
	num1 := 10
	num2 := 5
	str1 := "Hello"
	str2 := "World"

	fmt.Println("Equal to:")
	fmt.Println(num1 == num2)
	fmt.Println(str1 == str2)

	fmt.Println("Not equal to:")
	fmt.Println(num1 != num2)
	fmt.Println(str1 != str2)

	fmt.Println("Greater than:")
	fmt.Println(num1 > num2)

	fmt.Println("Less than:")
	fmt.Println(num1 < num2)

	fmt.Println("Greater than or equal to:")
	fmt.Println(num1 >= num2)

	fmt.Println("Less than or equal to:")
	fmt.Println(num1 <= num2)
}

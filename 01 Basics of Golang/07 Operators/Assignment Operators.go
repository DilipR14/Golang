package main

import "fmt"

func main() {
	num := 10

	num += 5
	fmt.Println("Addition assignment:", num)

	num -= 3
	fmt.Println("Subtraction assignment:", num)

	num *= 2
	fmt.Println("Multiplication assignment:", num)

	num /= 4
	fmt.Println("Division assignment:", num)

	num %= 3
	fmt.Println("Remainder assignment:", num)

	num &= 0xF
	fmt.Println("Bitwise AND assignment:", num)

	num |= 0x8
	fmt.Println("Bitwise OR assignment:", num)

	num ^= 0b1010
	fmt.Println("Bitwise XOR assignment:", num)

	num <<= 2
	fmt.Println("Left shift assignment:", num)

	num >>= 3
	fmt.Println("Right shift assignment:", num)
}

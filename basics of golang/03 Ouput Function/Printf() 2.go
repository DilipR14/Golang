package main

import (
	"fmt"
)

func main() {
	number := 42

	fmt.Printf("Decimal: %d\n", number)

	fmt.Printf("Binary: %b\n", number)

	fmt.Printf("Hexadecimal (lowercase): %x\n", number)

	fmt.Printf("Hexadecimal (uppercase): %X\n", number)
}

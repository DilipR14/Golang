package main

import (
	"fmt"
)

func main() {
	var x float64 = 1.7e+308
	fmt.Printf("Type: %T, value: %v", x, x)
}

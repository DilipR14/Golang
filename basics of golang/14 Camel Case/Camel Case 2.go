package main

import (
	"fmt"
)

func CalculateTotalpen(numitem int, unitpen float64) float64 {
	return float64(numitem) * unitpen
}
func main() {
	pen := CalculateTotalpen(5, 69.5)
	fmt.Println("Totalpen:", pen)
}

package main

import (
	"fmt"
)

func myFunction(x int, y int) (result int) {
	result = x + y
	return
}

func main() {
	fmt.Println(myFunction(400, 20))
}

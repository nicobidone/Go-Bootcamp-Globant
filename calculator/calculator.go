package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nicob/calculator/operations"
)

func main() {

	a, e := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[3])
	if e == nil {
		switch os := os.Args[2]; os {
		case "+":
			fmt.Println(operations.Add(a, b))
		case "-":
			fmt.Println(operations.Sub(a, b))
		case "*":
			fmt.Println(operations.Mul(a, b))
		case "/":
			fmt.Println(operations.Div(a, b))
		}
	}
	// Example 3 + 5
	// Return 8
}

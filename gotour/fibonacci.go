package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	sum := 1
	sumant := 0
	aux := 0
	return func(x int) int {
		if x == 0 {
			return 0
		}

		fmt.Print(sum, " ", sumant, " -> ")
		aux = sum
		sum += sumant
		sumant = aux

		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
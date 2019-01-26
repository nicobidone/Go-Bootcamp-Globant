package main

import (
	"fmt"
)

func whatType(v interface{}) {
	fmt.Printf("type of %v: %T\n", v, v)
}

type str struct{}

func main() {

	/*
		_, err := strconv.Atoi(os.Args[1])
		if err == nil {
			fmt.Println("integer")
			return
		}
		b := os.Args[1]
		if b == "true" || b == "false" {
			fmt.Println("bool")
			return
		}
		_, err2 := strconv.ParseFloat(os.Args[1], 64)
		if err2 == nil {
			fmt.Println("float64")
			return
		}
		fmt.Println("string")
	*/
	s := str{}
	whatType(true)
	whatType(s)
	whatType(3)
	whatType("hi")
	whatType(5.3)

}

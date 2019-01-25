package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

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

	/*
		wordPtr := flag.String("word", "foo", "a string")
		numbPtr := flag.Int("numb", 42, "an int")
		boolPtr := flag.Bool("fork", false, "a bool")
		floatPtr := flag.Float64("float64", 0.1, "a float")
		var svar string
		flag.StringVar(&svar, "svar", "bar", "a string var")
		flag.Parse()
		fmt.Println("word:", *wordPtr)
		fmt.Println("numb:", *numbPtr)
		fmt.Println("fork:", *boolPtr)
		fmt.Println("lefloat", *floatPtr)
		fmt.Println("svar:", svar)
		fmt.Println("tail:", flag.Args())
	*/
}

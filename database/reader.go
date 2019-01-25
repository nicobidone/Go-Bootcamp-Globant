package main

import (
	"fmt"
	"os"
)

func main() {

	switch oe := os.Args[1]; oe {

	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		for i := 2; i < len(os.Args); i++ {
			fmt.Print(os.Args[i], " ")
		}
	}
}

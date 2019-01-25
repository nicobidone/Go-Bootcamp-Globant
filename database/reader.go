package main

import (
	"fmt"
	"os"
)

func main() {

	switch oe := os.Args[1]; oe {
	case "PostgreSQL":
		fmt.Println("PostgreSQL")
	case "MongoDB":
		fmt.Println("MongoDB")
	default:
		for i := 2; i < len(os.Args); i++ {
			fmt.Print(os.Args[i], " ")
		}
	}

	//Example: MongoDB otro texto de prueba
}

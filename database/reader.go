package main

import (
	"fmt"
	"os"

	"github.com/nicob/database/selecteds"
)

func main() {

	switch oe := os.Args[1]; oe {
	case "PostgreSQL":
		fmt.Println(selecteds.PostgreSQL())
	case "MongoDB":
		fmt.Println(selecteds.MongoDB())
	default:
		for i := 2; i < len(os.Args); i++ {
			fmt.Print(os.Args[i], " ")
		}
	}

	//Example: MongoDB otro texto de prueba
}

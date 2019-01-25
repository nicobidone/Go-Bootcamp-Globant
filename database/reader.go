package main

import (
	"fmt"

	"github.com/nicob/database/selecteds"
)

func main() {

	/*
		   	switch oe := os.Args[1]; oe {
		   	case "SELECT":
		   		fmt.Println(selecteds.Pused())
		   	case "db.inventory.find(":
		   		fmt.Println(selecteds.Mused())
		   	default:
				for i := 2; i < len(os.Args); i++ {
					fmt.Print(os.Args[i], " ")
				}
				fmt.Println("Syntax not recognised")
		   }
		//Example: 	"db.inventory.find(" "{} )"
		//Return:	"You used MongoDB DB"
		//Example:	SELECT * FROM inventory
		//Return:	"You used PostgreSQL DB"
	*/

	//var x selecteds.Doer
	//x = &selecteds.MongoDB{}

	dbs := []selecteds.Doer{selecteds.PostgreSQL{}, selecteds.MongoDB{}}
	for _, db := range dbs {
		fmt.Println(db.Decir())
	}

}

package main

import (
	"fmt"

	"github.com/nicob/db/thedata"

	"github.com/nicob/db/themap"
)

func main() {
	x := themap.NewMap()
	b := *thedata.NewContainer("a", "b")
	x.Create("a", b)
	fmt.Println(x.Read("a"))
}

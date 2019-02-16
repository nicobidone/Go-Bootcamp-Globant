package main

import (
	"fmt"

	"github.com/nicob/theRESTservice/thereader"
)

func main() {
	url := "http://challenge.getsandbox.com/articles"
	fmt.Println(*thereader.GetJson(url))
}

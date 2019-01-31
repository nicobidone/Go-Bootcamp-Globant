package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// SafeCounter struct
type SafeCounter struct {
	v int
	//ch  chan string
	mux sync.Mutex
}

func (c *SafeCounter) sender(msg chan<- string, pun **bool) {
	for i := 0; i < 100; i++ {
		//fmt.Println("sender ", &(*pun), "->")
		//fmt.Println(**pun)
		if !(**pun) {
			return
		}
		c.mux.Lock()
		time.Sleep(1 * time.Second)
		c.v++
		msg <- "Sended " + strconv.Itoa(c.v)
		c.mux.Unlock()

	}
}

/*
func sender(msg chan<- string) {
	for i := 0; i < 100; i++ {

		time.Sleep(1 * time.Second)

		msg <- "Sended " + strconv.Itoa(i)
		//c.ch <- "Sended " + strconv.Itoa(c.v)

	}
}*/

func printer(msg <-chan string, pun **bool) { //(c *SafeCounter) printer(msg <-chan string) {
	//c.mux.Lock()
	m := <-msg
	for i := 0; i < 100; i++ {
		//fmt.Print("printer ", &(*pun), " -> ")
		fmt.Println("received", m)
		m = <-msg
	}
	//c.mux.Lock()

}

func stopper(pun **bool) {
	time.Sleep(5 * time.Second)
	b := false
	*pun = &b
	fmt.Println("Stopped!")
	//fmt.Println("stopper ", &pun)
}

func main() {

	ch1 := make(chan string)
	c := SafeCounter{v: 0} //, ch: ch1}

	var someBool *bool
	b := true
	someBool = &b
	fmt.Println("someBool ", &someBool, " -> ")
	go stopper(&someBool)

	//for i := 0; i < 5; i++ { // Completa los envios
	go c.sender(ch1, &someBool)
	//go sender(ch1)
	//}
	//time.Sleep(5 * time.Second) // no cambia nada

	//go c.printer(ch1)
	go printer(ch1, &someBool)

	time.Sleep(10 * time.Second)
	//fmt.Println(c.v)
}

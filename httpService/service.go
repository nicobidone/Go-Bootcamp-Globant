package main

import (
	"fmt"
	"log"
	"net/http"
)

// Name holds a name and implements http.Handler.
type Name string

func (n Name) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if _, err := w.Write([]byte(fmt.Sprintf("Hello %s", n))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main2() {
	n := Name("Globant")
	http.Handle("/", n)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

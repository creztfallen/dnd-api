package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	}

	http.HandleFunc("/hello", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

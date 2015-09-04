package main

import (
	"fmt"
	"net/http"
)

func main() {
	db()
	http.HandleFunc("/", handler)

	// create a listener on 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("listener error")
		panic(err)
	}
}

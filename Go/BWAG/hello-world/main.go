package main

import (
	"fmt"
	"net/http"
)

func main() {
	//fmt.Println("Hello, world")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, World! this is Trevor!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d:", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}

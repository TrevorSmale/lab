package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve static files from the /public directory
	http.Handle("/", http.FileServer(http.Dir("./public")))

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Define an API endpoint at "/api"
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is an API endpoint")
	})

	fmt.Println("Server is running on port 8081")
	http.ListenAndServe(":8081", nil)
}

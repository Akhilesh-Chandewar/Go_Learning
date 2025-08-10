package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from users")
	})

	port := 8080
	fmt.Printf("Server is running on port %d\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

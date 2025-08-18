package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3000"
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, "Hello World")
		w.Write([]byte("Hello World"))
	})
	fmt.Printf("Starting server on port %s\n", port)
	err := http.ListenAndServe(port , nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
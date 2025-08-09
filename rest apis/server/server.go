package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w , "Hello from server")
	})


	// const serverAddr string = "127.0.0.1:8080"
	const port string = ":8080"

	fmt.Println("Server is running on", port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println(err)
	}


}
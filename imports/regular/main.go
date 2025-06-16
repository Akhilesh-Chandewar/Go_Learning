package main

// regular imports
import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello from standard library")
	resp , err := http.Get("http://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

}

package main

// named imports
import (
	console "fmt"
	web "net/http"
)

func main() {
	console.Println("Hello from standard library")
	resp , err := web.Get("http://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		console.Println(err)
	}
	defer resp.Body.Close()
	console.Println("Response status:", resp.Status)
}


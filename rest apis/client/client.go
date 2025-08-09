package main

import (
	"fmt"
	"io"
	"net/http"
)

func main(){
	client := &http.Client{}

	response , err := client.Get("http://jsonplaceholder.typicode.com/posts/1")
	// response , err := client.Get("https://swapi.dev/api/people/1")

	if err != nil {
		fmt.Println("Error while making get" , err)
		panic(err)
	}
	defer response.Body.Close()

	body , err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error while reading body" , err)
		panic(err)
	}
	fmt.Println("Response body" , string(body))

}
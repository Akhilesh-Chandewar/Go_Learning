package main

import "fmt"

func main() {
	num := 42
	fmt.Printf("%05d\n", num)

	message := "Hello, World!"
	fmt.Printf("|%20s|\n", message)
	fmt.Printf("|%-20s|\n", message)

	message1 := "Go \n Programming"
	message2 := `Go \n Programming`
	fmt.Println(message1)
	fmt.Println(message2)
}
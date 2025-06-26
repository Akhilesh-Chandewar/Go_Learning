package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Deffered statement")
	fmt.Println("Starting the main fuction")
	os.Exit(1)
	fmt.Println("Ending the main fuction")
}
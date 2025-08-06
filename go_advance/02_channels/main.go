package main

import (
	"fmt"
	"time"
)

func main() {
	// variable := make(chan type)
	// variable <- 1
	// fmt.Println(<-variable)

	greet := make(chan string)
	greetString := "Hello"
	go func() {
		greet <- greetString // blocking operation 
		greet <- "World"

		for _ , e := range "aeiou" {
			greet <- "Vowels " + string(e)
		}
	}()
	go func ()  {
		receiver := <-greet
		fmt.Println(receiver)
		receiver = <-greet
		fmt.Println(receiver)
		for range 5 {
			receiver = <-greet
			fmt.Println(receiver)
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("End of program")
}
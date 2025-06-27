package main

import "fmt"

func main() {
	var age int = 20
	if age > 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	temperature := 25
	if temperature > 30 {
		fmt.Println("It's a hot day")
	} else if temperature > 20 {
		fmt.Println("It's a nice day")
	} else {
		fmt.Println("It's a cold day")
	}

	num := 18 
	if num % 2 == 0 {
		if num % 3 == 0 {
			fmt.Println("Number is divisible by 2 and 3")
		} else {
			fmt.Println("Number is divisible by 2 but not 3")
		}
	} else {
		fmt.Println("Number is not divisible by 2")
	}
}
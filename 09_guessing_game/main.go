package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	target := r.Intn(100) + 1

	fmt.Println("Welcome to the Guessing Game")
	fmt.Println("The target number is between 1 and 100")
	fmt.Println("Try and guess the number")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		if guess == target {
			fmt.Println("You guessed it! Congratulations!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
	}
}

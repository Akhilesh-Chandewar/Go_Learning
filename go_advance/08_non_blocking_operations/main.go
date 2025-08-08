package main

import (
	"fmt"
	"time"
)

func main() {
	ch  := make(chan int)
	// Non blocking receive operations
	select {
	case msg := <- ch:
		fmt.Println("Received", msg)
	default:
		fmt.Println("No message received")
	}

	//Non blocking send operation
	select {
	case ch <- 1:
		fmt.Println("Sent")
	default:
		fmt.Println("No message sent")
	}

	data := make(chan int) 
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Received", d)
			case <-quit:
				fmt.Println("Quit")
				return
			default:
				fmt.Println("No message received")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}
	quit <- true
}
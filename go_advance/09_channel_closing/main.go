package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple closing channel
	ch := make(chan int)
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
	for value := range ch {
		fmt.Println("Received: ", value)
	}

	// receiving from closed channel
	ch2 := make(chan int)
	close(ch2)
	value , ok := <- ch2
	if !ok {
		fmt.Println("Channel closed")
	} else {
		fmt.Println("Received: ", value)
	}

	// range over closed channel
	ch3 := make(chan int)
	go func(){
		for i := range 5 {
			ch3 <- i
		}
		close(ch3)
	}()
	for value := range ch3 {
		fmt.Println("Received: ", value)
	}

	// run time panic
	ch4 := make(chan int)
	go func(){
		close(ch4)
		// close(ch4)
	}()
	time.Sleep(1 * time.Second)

	ch5 := make(chan int)
	ch6 := make(chan int)
	go producer(ch5)
	go filter(ch5, ch6)
	for value := range ch6 {
		fmt.Println("Received: ", value)
	}
}

func producer(ch chan <- int){
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

func filter(in <- chan int, out chan <- int){
	for value := range in {
		if value % 2 == 0 {
			out <- value
		}
	}
	close(out)
}
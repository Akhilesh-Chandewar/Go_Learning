package main

import "fmt"

func main() {
	process(10)
	fmt.Println("main")
}

func process (i int) {
	defer fmt.Println(i)
	defer fmt.Println("Process")
	defer fmt.Println("Middle")
	defer fmt.Println("End")
	fmt.Println("Start")
	i++
	fmt.Println(i)
}
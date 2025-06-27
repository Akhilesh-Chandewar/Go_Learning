package main

import "fmt"

func main() {
	// panic(interface{})

	// process(1)
	process(-1)
}

func process(i int){
	defer fmt.Println("Deffered 1")
	defer fmt.Println("Deffered 2")
	if i < 0{
		fmt.Println("Before panic")
		panic("i should be non negative number")
		// defer fmt.Println("After panic")
	}
	fmt.Println(i)
}
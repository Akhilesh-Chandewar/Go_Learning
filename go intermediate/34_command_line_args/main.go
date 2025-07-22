package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("command:", os.Args[0])
	for i, arg := range os.Args[1:] {
		fmt.Println("arg:", i , arg)
	}

	//define flags
	var name string
	var age int

	flag.StringVar(&name, "name", "John Doe", "Name of the person")
	flag.IntVar(&age, "age", 30, "Age of the person")

	flag.Parse()
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}
package main

import "fmt"

func main() {
	fmt.Println("*****")
	process()
	fmt.Println("*****")
}

func process(){
	defer func ()  {
		r := recover()
		if r != nil {
			fmt.Println("Recover" , r)
		} 
	}()
	fmt.Printf("Start process \n")
	panic("Something went wrong")
	fmt.Printf("End process \n")
}
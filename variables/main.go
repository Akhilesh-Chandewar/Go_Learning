package main

import "fmt"

// middleName := "somemillename"
var middleNameOutside string = "somemillenameoutside" //accessable only in same package

func main() {
	var age int  //uninitialized
	fmt.Println(age)

	var name string = "John"
	fmt.Println(name) 

	var name1 = "Jane"
	fmt.Println(name1)

	count := 10
	fmt.Println(count)

	lastName := "Smith"
	fmt.Println(lastName)

	//Default values
	var number int
	var str string
	var isTrue bool
	fmt.Println("default", number, str, isTrue)
	//pointer slices maps functions and structs => nil

	
	//Scopes
	// blocked scope
	printName()
	// fmt.Println(firtstName)
	middleName := "somemillename"
	fmt.Println(middleName)
	fmt.Println(middleNameOutside)
	middleName = "someothername"
	fmt.Println(middleName)
	middleNameOutside = "someothernameoutside"
	fmt.Println(middleNameOutside)
}

func printName()  {
	firtstName := "John"
	fmt.Println(firtstName)		
}
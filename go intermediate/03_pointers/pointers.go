package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")
	// var name *type
	var a int 
	var ptr *int
	ptr = &a
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of ptr:", ptr)
	fmt.Println("Value of *ptr:", *ptr) 

	var ptr2 *int
	fmt.Println("Value of ptr2:", ptr2)
	ptr2 = new(int)
	fmt.Println("Value of ptr2 after new:", ptr2)
	*ptr2 = 42
	fmt.Println("Value of *ptr2:", *ptr2)

	x := 10
	fmt.Println("Value of x:", x)
	changeValue(&x)
	fmt.Println("Value of x after change:", x)

	y := &x
	fmt.Println("Value of y:", y)
	fmt.Println("Value of *y:", *y)

}

func changeValue(x *int) {
	*x = *x + 100
}
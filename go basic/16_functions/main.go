package main

import "fmt"

func main() {
	// func <name> (params) returnType {
	// 	//code
	// return value
	// }

	sum := add(1, 2)
	fmt.Println(sum)

	func() {
		fmt.Println("Anonymous function")
	}()

	greet := func(name string) {
		fmt.Println("Hello", name)
	}
	greet("John")

	operation := add
	fmt.Println(operation(1, 2))

	result := applyOperation(1, 2, add)
	fmt.Println(result)

	multiplier := createMultiplier(2)
	fmt.Println(multiplier(5))
}

func add(a int, b int) int {
	return a + b
}
	
func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func createMultiplier(factor int) func(int) int {
	return func(value int) int {
		return value * factor
	}
}
package main

import (
	"fmt"
	"slices"
)

func main() {
	// Declare an empty slice of int
	var numbers []int
	fmt.Println(numbers) // Output: []

	// Declare and initialize a slice using var
	var numbers1 = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers1) // Output: [1 2 3 4 5]

	// Declare and initialize a slice using :=
	numbers2 := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers2) // Output: [1 2 3 4 5]

	// Create a slice using make with length 5
	slice := make([]int, 5)
	fmt.Println(slice) // Output: [0 0 0 0 0]

	// Create an array and slice a portion of it
	array := [5]int{1, 2, 3, 4, 5}
	slc := array[1:4]
	fmt.Println(slc) // Output: [2 3 4]

	slc = append(slc, 6 , 7)
	fmt.Println(slc) // Output: [2 3 4 6 7]

	slcCopy := make([]int , len(slc))
	copy(slcCopy , slc)
	fmt.Println(slcCopy)

	var nilSlice []int
	fmt.Println(nilSlice)

	for idx , value := range slc {
		fmt.Println("index" , idx , "value" , value)
	}

	fmt.Println("len" , len(slc)) //
	fmt.Println("cap" , cap(slc))

	fmt.Println("Value at index 3" , slc[3])
	// slc[3] = 100
	fmt.Println("Value at index 3" , slc[3])

	if slices.Contains(slc , 4) {
		fmt.Println("4 is present in the slice")
	}

	if slices.Equal(slc , slcCopy) {
		fmt.Println("slc and slcCopy are equal")
	}

	twoDimSlice := [][]int{{1,2,3} , {4,5,6}}
	fmt.Println(twoDimSlice)

	twoDim := make([][]int, 3)
	for i := range 3 {
		twoDim[i] = make([]int, 3)
		for j := range i {
			twoDim[i][j] = i + j
		}
	}

	// Print the 2D slice
	fmt.Println(twoDim)

	slice1 := slice[1:3]
	fmt.Println(slice1 , cap(slice1), len(slice1))
}

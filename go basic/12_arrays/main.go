package main

import "fmt"

func main() {
	// var arrayName [size] type

	var numbers [5] int
	fmt.Println(numbers) // [0 0 0 0 0]

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	fmt.Println(numbers)
	
	fruits := [3]string{"apple" , "banana" , "orange"}
	fmt.Println(fruits)
	
	originalArray := [3]int{1,2,3}
	copiedArray := originalArray
	copiedArray[0] = 10 //shallow copy
	fmt.Println(originalArray)
	fmt.Println(copiedArray) 

	for i:=range len(numbers) {
		fmt.Println(numbers[i])
	}

	for key , value := range numbers {
		fmt.Println("key :" , key , "value :" , value)
	}

	_ , numbers2 := someFunction()
	fmt.Println(numbers2)

	fmt.Println("length of numbers" , len(numbers)) // length of numbers 5

	array1 := [3]int{1,2,3}
	array2 := [3]int{4,5,6}
	array3 := [3]int{1,2,3}
	fmt.Println(array1 == array2) // false
	fmt.Println(array1 == array3) // true

	var matrix[3][3] int = [3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println(matrix)
	fmt.Println(matrix[1][1])

	original := [3] int {1,2,3}
	var copy *[3]int = &original
	copy[0] = 10
	fmt.Println(original) // [10 2 3]
	fmt.Println(copy) // &[10 2 3]
	fmt.Println(*copy) // [10 2 3]

}

func someFunction () (int ,int){
	return 1,2
}
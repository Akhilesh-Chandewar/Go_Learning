package main

import "fmt"

func main() {

	// ...Ellipsis
	// func name(param1 type1 , param2 type2 , param3 ...type3) returnType {
	// 	//code
	// 	return value
	// }

	add := func( str string , values ...int , ) (string , int) {
		total := 0
		for _, value  := range values {
			total += value
		}
		return "result"+str , total
	}
	fmt.Println(add("string" ,1, 2, 3, 4, 5))

	numbers := []int{1,2,3,4,5} 
	fmt.Println(add("string" ,numbers...))
}
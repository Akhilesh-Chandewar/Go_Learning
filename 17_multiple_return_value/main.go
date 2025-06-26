package main

import (
	"errors"
	"fmt"
)

func main() {
	// func name(param1 type1 , param2 type2) (return1 type1, return2 type2) {
	// 	//code
	// 	return value1, value2
	// }

	quotient, remainder := divide(10, 3)
	fmt.Println(quotient, remainder)

	message, err := compare(10, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message)
	}
}

func divide(a int, b int)(quotient int, remainder int){
	quotient = a / b
	remainder = a % b
	return 
}

func compare(a int, b int) (string, error){
	if(a > b){
		return "a is greater than b", nil
	} else if(a < b){
		return "a is less than b", nil
	} else {
		return "" , errors.New("Unable to compare")
	}

}
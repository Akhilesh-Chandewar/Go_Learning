package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	// fmt.Println(i)
	// }

	// numbers := []int{1,2,3,4,5}
	// // for index, value := range numbers {
	// // 	// fmt.Printf("index: %d : value: %d \n", index, value)
	// // }

	// for _ , value := range "Hello World" {
	// 	fmt.Printf("%c \n" , value)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i % 2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// rows := 5
	// for i:=1; i<=rows; i++ {
	// 	for j:=1; j<=rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k:=1; k<=2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// for i:= range 10 { //inclusive 0 exclusive 10
	// 	fmt.Println(10 - i)
	// }

	for i:= range "Hello World" {
		fmt.Printf("%c \n" , "Hello World"[i])
	}
}
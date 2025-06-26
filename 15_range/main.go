package main

import "fmt"

func main() {
	// az := "azAZ09"
	// for i, v := range az {
	// 	fmt.Println(i, v)
	// 	// fmt.Printf("Index %d : Rumes %c \n" , i, v)
	// }
	
	message := "Hello world"
	for i, v := range message {
		// fmt.Println(i, v)
		fmt.Printf("Index %d : Rumes %c \n" , i, v)
	}

	
}
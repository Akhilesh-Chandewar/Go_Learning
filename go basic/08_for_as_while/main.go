package main

import "fmt"

func main() {
	// i:=0
	// for i<=5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// //infinite loop
	// for {
	// 	fmt.Println("Hello")
	// }

	// sum := 0
	// for {
	// 	sum += 10
	// 	fmt.Println(sum)
	// 	if sum >= 50 {
	// 		break
	// 	}
	// }
	
	i:=1;
	for i<=10{
		if(i % 2 == 0){
			i++
			continue
		}
		fmt.Println(i)
		i++
	}
}
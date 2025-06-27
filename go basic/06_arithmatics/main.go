package main

import (
	"fmt"
	"math"
)

func main() {
	var a , b int = 10, 3
	var res int 

	res = a + b
	fmt.Println("addition" , res)

	res = a - b
	fmt.Println("subtraction" , res)

	res = a * b
	fmt.Println("multiplication" , res)

	res = a / b
	fmt.Println("division" , res)

	res = a % b
	fmt.Println("remainder" , res)

	const pi float32 = 22/7.0
	fmt.Println("pi" , pi) 

	var maxi int64 = 9223372036854775807
	var mini int64 = -9223372036854775808
	fmt.Println("maxi" , maxi)
	fmt.Println("mini" , mini)
	maxi = maxi + 1
	fmt.Println("maxi" , maxi) //obverflow with signed numbers
	mini = mini - 1
	fmt.Println("mini" , mini)

	var maxu uint8 = 255
	var minu uint8 = 0
	fmt.Println("maxu" , maxu)
	fmt.Println("minu" , minu)
	maxu = maxu + 1
	fmt.Println("maxu" , maxu) //obverflow with unsigned numbers
	minu = minu - 1
	fmt.Println("minu" , minu)

	var smallFLoat float64 = 1.0e-323
	fmt.Println("smallFLoat" , smallFLoat)
	smallFLoat = smallFLoat / math.MaxFloat64
	fmt.Println("smallFLoat" , smallFLoat)

}

package main

import "fmt"

func main() {
	fmt.Println("closures")
	sequence := adder()
	fmt.Println("first call:", sequence())
	fmt.Println("second call:", sequence())
	fmt.Println("third call:", sequence())

	fmt.Println(">>>>>>>>>")
	sequence2 := adder()
	fmt.Println("first call:", sequence2())
	fmt.Println("second call:", sequence2())
	fmt.Println("third call:", sequence2())

	fmt.Println(">>>>>>>>>")

	substaractor := func () func(int) int {
		countdown := 100
		return func(i int) int {
			countdown -= i
			fmt.Println("current value of countdown:", countdown)
			return countdown
		}
	}

	fmt.Println("subtractor")
	substractor := substaractor()
	fmt.Println("first call:", substractor(10))
	fmt.Println("second call:", substractor(20))
	fmt.Println("third call:", substractor(30))
}

func adder() func() int{
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("current value of i:", i)
		return i
	}
}
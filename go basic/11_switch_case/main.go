package main

import "fmt"

func main() {

	// fruit:= "apple"
	// switch fruit {
	// case "apple":
	// 	fmt.Println("apple")
	// case "banana":
	// 	fmt.Println("banana")
	// default:
	// 	fmt.Println("unknown")
	// }

	// day:= "Monday"
	// switch day {
	// case "Monday" , "Tuesday" , "Wednesday" , "Thursday" , "Friday":
	// 	fmt.Println("weekday")
	// case "Saturday" , "Sunday":
	// 	fmt.Println("weekend")
	// default:
	// 	fmt.Println("unknown")
	// }

	// number := 15
	// switch {
	// case number <= 10:
	// 	fmt.Println("less than or equal to 10")
	// case number <= 20:
	// 	fmt.Println("less than or equal to 20")
	// default:
	// 	fmt.Println("greater than 20")
	// }

	// num:= 2
	// switch {
	// case num >1 :
	// 		fmt.Println("num is greater than 1")
	// 		fallthrough
	// case num == 2:
	// 		fmt.Println("num is 2")
	// default:
	// 		fmt.Println("num is less than 1")
	// }

	checkType(10)
	checkType("hello")
	checkType(true)
	checkType(3.14)
}

func checkType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("unknown")
	}
}
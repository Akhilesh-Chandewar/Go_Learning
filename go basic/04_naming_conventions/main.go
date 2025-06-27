package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//PascalCase : Struct , Interfaces , Enums
	//Eg. CalculateArea, UserInfo, NewHttpRequest

	//snake_case : 
	//Eg. user_id, first_name, http_request

	//UPPERCASE
	//Constants

	//mixedCase
	//Eg. javaScript , isValid

	const MAXRETRIES = 5
	var employeeID = 1001
	fmt.Println(employeeID)
}

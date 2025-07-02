package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	// EmployeeInfo Person //named fiel
	Person // Embedded struct // Anonymous field
	ID     string
	Salary float64
}

func (p Person) introduce() string {
	return fmt.Sprintf("My name is %s and I am %d years old.", p.Name, p.Age)
}

func (e Employee) introduce() string {
	return fmt.Sprintf("I am %s, an employee with ID %s and a salary of %.2f.", e.Name, e.ID, e.Salary)
}

func main() {

	emp1 := Employee{
		Person: Person{
			Name: "John Doe",
			Age:  30,
		},
		ID:     "12345",
		Salary: 50000.0,
	}
	fmt.Println(emp1)
	fmt.Println("Employee Name:", emp1.Name) // Accessing embedded struct field
	fmt.Println("Employee Age:", emp1.Age)   // Accessing embedded struct field
	fmt.Println("Employee ID:", emp1.ID)
	fmt.Println("Employee Salary:", emp1.Salary)
	fmt.Println(emp1.introduce()) // Calling method from embedded struct
}
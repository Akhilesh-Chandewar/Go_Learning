package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city	string
	country	string
}

// func (p Person) fullName() string {
// 	return p.firstName + " " + p.lastName
// }

// func (p Person) increaseAge()  {
// 	p.age++
// }

// func (p *Person) setAge(age int) {
// 	p.age = age
// }

func main() {

	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "New York",
			country: "USA",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "123-456-7890",
			cell: "098-765-4321",
		},
	}
	// fmt.Println(p1)

	p2 := Person{
		firstName: "Jane",
		lastName:  "Doe",
	}
	fmt.Println(p2)
	// p2.address.city = "Los Angeles"
	// p2.address.country = "USA"
	// p2.PhoneHomeCell.cell = "555-555-5555"
	// p2.PhoneHomeCell.home = "555-555-1234"
	// fmt.Println("First Name:", p1.firstName)
	// fmt.Println("Last Name:", p1.lastName)
	// fmt.Println("Age:", p1.age)
	// fmt.Println("City:", p1.address.city)
	// fmt.Println("Country:", p1.address.country)
	// fmt.Println("First Name:", p2.firstName)
	// fmt.Println("Last Name:", p2.lastName)
	// fmt.Println("Age:", p2.age)
	// fmt.Println("City:", p2.address.city)
	// fmt.Println("Country:", p2.address.country)

	// Anonymous Struct
	// user := struct {
	// 	username string
	// 	email    string
	// }{
	// 	username: "johndoe",
	// 	email:    "dFZV9@example.com",
	// }
	// fmt.Println(user)

	// // Using a method on a struct
	// fmt.Println("Full Name:", p1.fullName())
	// p1.increaseAge()
	// fmt.Println("Age after increase:", p1.age)
	// p1.setAge(35)
	// fmt.Println("Age after set:", p1.age)
	// fmt.Println("person 1:", p1)

	// fmt.Println(p1 == p2)

	// p3 := Person{
	// 	firstName: "Jane",
	// 	lastName:  "Doe",
	// }
	// fmt.Println(p3)
	// fmt.Println(p2 == p3)

	//shallow copy
	p4 := p1
	fmt.Println("p4:", p4)
	p4.firstName = "Alice"
	fmt.Println("p1 after p4 change:", p1)
	fmt.Println("p4 after change:", p4)
}
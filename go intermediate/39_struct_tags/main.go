package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name" db:"first_name" xml:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"-"`
}

func main() {
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatalln("Error marshalling to JSON:", err)
		return
	}
	fmt.Println("JSON output:", string(jsonData))

	person2 := Person{
		FirstName: "Jane",
		Age:       25,
	}
	jsonData2, err := json.Marshal(person2)
	if err != nil {
		log.Fatalln("Error marshalling to JSON:", err)
		return
	}
	fmt.Println("JSON output:", string(jsonData2))
}
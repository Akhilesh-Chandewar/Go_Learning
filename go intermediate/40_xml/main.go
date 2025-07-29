package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
	Address Address  `xml:"address"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	ISBN	string   `xml:"isbn,attr"`
	Title   string   `xml:"title,attr"`
	Author  string   `xml:"author"`
}

func main() {
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "jane@fakemail.com",
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}
	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error marshalling to XML:", err)
	}
	fmt.Println("XML output:", string(xmlData))

	xmlDataIndented, err := xml.MarshalIndent(person, "", " ")
	if err != nil {		
		log.Fatalln("Error marshalling to XML with indentation:", err)
	}
	fmt.Println("Indented XML output:")
	fmt.Println(string(xmlDataIndented))

	// Unmarshal XML back to struct
	var personUnmarshalled Person
	err = xml.Unmarshal(xmlData, &personUnmarshalled)
	if err != nil {
		log.Fatalln("Error unmarshalling XML:", err)
	}
	fmt.Println("Unmarshalled Person struct:", personUnmarshalled)

	fmt.Println("Local" , personUnmarshalled.XMLName.Local)
	fmt.Println("Space" , personUnmarshalled.XMLName.Space)

	book := Book{
		ISBN:  "978-3-16-148410-0",
		Title: "Go Programming",
		Author:  "John Doe",
	}
	xmlBookData, err := xml.MarshalIndent(book , "", "  ")
	if err != nil {
		log.Fatalln("Error marshalling to XML:", err)
	}
	fmt.Println("Book output:", string(xmlBookData))
	
	// Unmarshal XML back to struct
	var bookUnmarshalled Book
	err = xml.Unmarshal(xmlBookData, &bookUnmarshalled)
	if err != nil {
		log.Fatalln("Error unmarshalling XML:", err)
	}
	fmt.Println("Unmarshalled Book struct:", bookUnmarshalled)
}
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hello, \nGo!" 
	rawMessage := `Hello\nGo!`
	message2 := "Hello, \rGo!"

	fmt.Println(message)       
	fmt.Println(rawMessage)
	fmt.Println(message2) 

	fmt.Println("Length of message:", len(message))
	fmt.Println("Length of rawMessage:", len(rawMessage))
	fmt.Println("Length of message2:", len(message2))

	fmt.Println("First character of message:", (message[0])) // ASCII value of 'H'
	fmt.Println(`\n :`, message[5]) // ASCII value of newline character
	fmt.Println("Ascii for space:", message[6]) // ASCII value of space character


	greeting := "Hello"
	name := "John"
	fmt.Println(greeting + " " + name)
	fmt.Printf("%s %s \n", greeting, name)

	str1 := "banana"
	str2 := "bbnana"
	// String comparison all ascii values
	fmt.Println(str1 < str2)
	fmt.Println(str1 > str2)
	fmt.Println(str1 == str2)
	fmt.Println(str1 != str2)

	for i, char := range message {
		fmt.Printf("Character at index %d: %c\n", i, char)
	}

	for _, char := range rawMessage {
		fmt.Printf("Character in rawMessage: %v %x\n", char, char)
	}

	fmt.Println("Length of message in runes:", len([]rune(message)))
	fmt.Println("Length of rawMessage in runes:", len([]rune(rawMessage)))

	fmt.Println("Rune count in message:", utf8.RuneCountInString(message))
	fmt.Println("Rune count in rawMessage:", utf8.RuneCountInString(rawMessage))

	greetingWithName := greeting + " " + name
	fmt.Println("Greeting with name:", greetingWithName)

	var ch rune = 'a'
	fmt.Println("Character:", ch)
	fmt.Printf("Type of ch: %T\n", ch)
	chstr:= string(ch)
	fmt.Printf("Type of ch: %T\n", chstr)

	
}
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello 123 World!"
	fmt.Println("Original String:", str)
	fmt.Println("Concatenated String:", str+" How are you?")
	fmt.Println("Length of String:", len(str))
	fmt.Println("First Character:", string(str[0]))
	fmt.Println("Substring (0-5):", str[0:5])
	fmt.Println("Sunstring" , str[3:])
	fmt.Println("Sunstring" , str[:9])

	number := 12345
	strNumber := strconv.Itoa(number)
	fmt.Println("Integer to String:", strNumber)
	fmt.Println("strNumber Length:", len(strNumber))
	
	fruites := "apple, banana, cherry"
	parts := strings.Split(fruites, ", ")
	fmt.Println("Split String:", parts)

	countries := []string{"USA", "Canada", "Mexico"}
	countriesStr := strings.Join(countries, ", ")
	fmt.Println("Joined String:", countriesStr)

	fmt.Println("Contains 'World':", strings.Contains(str, "World"))
	fmt.Println("Index of 'World':", strings.Index(str, "World"))
	fmt.Println("Replace 'World' with 'Gopher':", strings.Replace(str, "World", "Gopher", 1) ,">>" , str) //ReplaceAll
	fmt.Println("To Upper Case:", strings.ToUpper(str))
	fmt.Println("To Lower Case:", strings.ToLower(str))
	fmt.Println("Trimmed String:", strings.TrimSpace("   Hello, Gopher!   "))
	fmt.Println("Repeat String:", strings.Repeat("Go! ", 3))
	fmt.Println("Counter" , strings.Count(str, "o"))
	fmt.Println("HasPrefix 'Hello':", strings.HasPrefix(str, "Hello"))
	fmt.Println("HasSuffix 'World!':", strings.HasSuffix(str, "World!"))

	fmt.Println("String matching", regexp.MustCompile(`\d+`).FindAllString(str, -1))
	fmt.Println("Runes count::" , utf8.RuneCountInString(str))

	//string builder
	var sb strings.Builder
	sb.WriteString("Hello")
	sb.WriteString(" ")
	sb.WriteString("World!")
	fmt.Println("String Builder:", sb.String())

	sb.WriteRune('!')
	sb.WriteString(" How are you?")
	fmt.Println("Updated String Builder:", sb.String())

	fmt.Println("String Builder Length:", sb.Len())
	fmt.Println("String Builder Capacity:", sb.Cap())

	fmt.Println("String Builder Reset")
	sb.Reset()
	fmt.Println("After Reset String Builder:", sb.String())
	
	sb.WriteString("Starting fresh")
	fmt.Println("After Reset String Builder:", sb.String())
}
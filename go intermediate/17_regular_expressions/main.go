package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Hello, \"He said I am learning Go!\"")
	fmt.Println(`Hello, "He said I am learning Go!"`)

	re := regexp.MustCompile(`[a-zA-Z0-9._+\-%]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	// re := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	// re := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	fmt.Println(re.MatchString("akhil.chandewar00@gmail.com"))
	fmt.Println(re.MatchString("invalid-email@.com"))

	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	fmt.Println(re.MatchString("2023-10-05"))
	matches := re.FindStringSubmatch("202-10-05")
	if len(matches) > 0 {
		fmt.Println("Matched:", matches)
		fmt.Println("Year:", matches[1])
		fmt.Println("Month:", matches[2])
		fmt.Println("Day:", matches[3])
	} else {
		fmt.Println("No match found")
	}

	str := "The quick brown fox jumps over the lazy dog"
	re = regexp.MustCompile(`[aeiou]`)
	result := re.ReplaceAllString(str, "*")	
	fmt.Println("Original String:", str)
	fmt.Println("Modified String:", result)

	// i - case insensitive
	// m - multiline
	// s - dotall
	// U - unicode
	// g - global

	re = regexp.MustCompile(`(?i)quick`)
	fmt.Println(re.MatchString("Quick brown fox jumps over the lazy dog"))

	re = regexp.MustCompile(`(?m)^The`)
	fmt.Println(re.MatchString("The quick brown fox jumps over the lazy dog"))

	re = regexp.MustCompile(`(?s)quick.*dog`)
	fmt.Println(re.MatchString("The quick brown fox jumps over the lazy dog"))

	re = regexp.MustCompile(`(?U)quick.*dog`)
	fmt.Println(re.MatchString("The quick brown fox jumps over the lazy dog"))

	
}
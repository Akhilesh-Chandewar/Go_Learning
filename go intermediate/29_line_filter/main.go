package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	linenumber := 1 

	keyword := "Go"
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			updatedLine := strings.ReplaceAll(line, keyword, "****") // case-sensitive replace
			fmt.Printf("%d Filtered line: %v \n", linenumber , line)
			fmt.Printf("%d Updated line: %v \n", linenumber , updatedLine)
		}
		linenumber++;
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

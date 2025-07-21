package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func () {
		fmt.Println("Closing file")
		file.Close()
	}()
	// data := make([]byte, 20)
	// n, err := file.Read(data)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Read line:", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
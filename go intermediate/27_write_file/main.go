package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, World!\n")
	if err != nil {
		fmt.Println("Error writing string to file:", err)
		return
	}

	data := []byte("Hello, World! byte slice\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing byte data to file:", err)
		return
	}

	fmt.Println("File created and written successfully!")
}

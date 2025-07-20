package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello, buffio package!\n How are you"))
	data := make([]byte, 20)
	n , err := reader.Read(data)
	if err != nil {
		fmt.Println("Error:", err)	
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))


	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)	
		return
	}
	fmt.Println("Read line:", line)

	writer := bufio.NewWriter(os.Stdout)
	data = []byte("Hello, buffio package!\n How are you")
	n , err = writer.Write(data)
	if err != nil {
		fmt.Println("Error:", err)	
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error:", err)	
		return
	}
}
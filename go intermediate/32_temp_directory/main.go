package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
		// fmt.Println(err)
	}
}

func main() {
	tempFile, err  := os.CreateTemp("", "temp")
	checkError(err)
	fmt.Println(tempFile.Name())
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	tempDir , err := os.MkdirTemp("", "temp")
	checkError(err)
	fmt.Println(tempDir)
	defer os.RemoveAll(tempDir)
}
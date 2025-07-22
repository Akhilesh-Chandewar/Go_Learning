package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var content string

//go:embed basics/*
var basicsFolder embed.FS

func main() {
	fmt.Println("Embedded content:", content)

	// Read a specific file
	data, err := basicsFolder.ReadFile("basics/hello.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Embedded file content:", string(data))

	// Walk all files in basics folder
	err = fs.WalkDir(basicsFolder, "basics", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("WalkDir error:", err)
			return err
		}
		fmt.Println("Found embedded file:", path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

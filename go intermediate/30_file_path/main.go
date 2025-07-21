package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	relativePath := "./go_intermediate/30_file_path"
	fmt.Println(relativePath)
	absolutePath := "/home/username/go/src/github.com/username/go_intermediate/30_file_path"
	fmt.Println(absolutePath)

	//join the path
	joinedPath := filepath.Join("documents", "downloads" , "file.zip")
	fmt.Println(joinedPath)

	//clean the path
	cleanPath := filepath.Clean("/home/username/go/src/github.com/username/go_intermediate/30_file_path/../../")
	fmt.Println(cleanPath)

	dir , file := filepath.Split("/home/username/go/src/file.txt")
	fmt.Println(dir)
	fmt.Println(file)

	fmt.Println(filepath.Base("/home/username/go/src/file.txt"))
	fmt.Println(filepath.Ext("/home/username/go/src/file.txt"))

	fmt.Println("Is absolute path" , filepath.IsAbs(relativePath))
	fmt.Println("Is absolute path" , filepath.IsAbs(absolutePath))

	fmt.Println(strings.TrimSuffix(file , filepath.Ext(file)))

	rel , err := filepath.Rel("/home/username/go/src/github.com/username/go_intermediate/30_file_path", "/home/username/go/src/github.com/username/go_intermediate/30_file_path/30_file_path.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rel)

	absPath , err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(absPath)
} 
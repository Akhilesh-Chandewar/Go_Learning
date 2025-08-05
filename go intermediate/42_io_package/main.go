package main

import (
	"fmt"
	"io"
	"log"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n , err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading from reader" , err)
	}
	fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
}

func writeToWriter(w io.Writer, data string) {
	n, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing to writer", err)
	}
	fmt.Printf("Wrote %d bytes: %s\n", n, data)
}

func main() {
	fmt.Println("")
}
package main

import "fmt"

func main() {
	var a int = 42;
	b := int32(a) // Convert int to int32
	c := float64(b) // Convert int to float64
	// d := bool("correct")
	e := 3.14 // float64 literal
	f := int(e) // Convert float64 to int (truncates the decimal part)
	fmt.Println("f:", f ,"c:", c)

	g := "hello"
	h := []byte(g) // Convert string to byte slice
	fmt.Println("h:", h)

	i := []byte{255 , 72}
	j := string(i) // Convert byte slice to string
	fmt.Println("j:", j)

	
}
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte ("He~llo base64 coding!")
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded:", encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Decoded:", string(decoded))

	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe Encoded:", urlSafeEncoded)
	urlSafeDecoded, err := base64.URLEncoding.DecodeString(urlSafeEncoded)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("URL Safe Decoded:", string(urlSafeDecoded))
}
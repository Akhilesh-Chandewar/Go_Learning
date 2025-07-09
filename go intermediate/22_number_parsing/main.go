package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	strconvNum, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	fmt.Println("Converted string to int:", strconvNum , "Type:", fmt.Sprintf("%T", strconvNum))

	numistr, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		fmt.Println("Error converting string to int64:", err)
		return
	}
	fmt.Println("Converted string to int64:", numistr, "Type:", fmt.Sprintf("%T ", numistr))

	floatStr := "123.45"
	floatNum, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}
	fmt.Println("Converted string to float64:", floatNum, "Type:", fmt.Sprintf("%T", floatNum))

	binaryStr := "101010"
	binaryNum, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error converting binary string to int64:", err)
		return
	}
	fmt.Println("Converted binary string to int64:", binaryNum, "Type:", fmt.Sprintf("%T", binaryNum))

	hexStr := "1a2b3c"
	hexNum, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error converting hex string to int64:", err)
		return
	}
	fmt.Println("Converted hex string to int64:", hexNum, "Type:", fmt.Sprintf("%T", hexNum))
}
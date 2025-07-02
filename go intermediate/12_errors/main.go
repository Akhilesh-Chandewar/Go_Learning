package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64)(float64, error) {
	if x < 0 {
		return 0, errors.New("mathematical: square root of negative number")
	}
	return math.Sqrt(x), nil
}

func process(data []byte) error {
	if len(data)== 0 {
		return errors.New("data cannot be empty")
	}
	return nil
}

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("MyError: %s", e.Message)
}

func processError() error {
	return &MyError{Message: "Custom error occurred"}
}

func readConfig() error{
	return errors.New("failed to read config")
}

func readData() error {
	if err := readConfig(); err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func main() {
	result, err := sqrt(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root of 2 is", result)
	}

	result, err = sqrt(-2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root of -2 is", result)
	}

	data := []byte{}
	if err := process(data); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Data processed successfully")
	}

	err1 := processError()
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("No error occurred")
	}

	err2 := readData()
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Data read successfully")
	}
} 
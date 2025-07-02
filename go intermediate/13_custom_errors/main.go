package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	code   int
	message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}

func doSomething() error {
	return &CustomError{code: 404, message: "Resource not found"}
}

func doError() error {
	return errors.New("An unexpected error occurred")
}

func exError() error {
	err := doError()
	if err != nil {
		return &CustomError{code: 500, message: err.Error()}
	}
	return nil
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Operation successful")
	}

	err = exError()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No error occurred")
	}
}
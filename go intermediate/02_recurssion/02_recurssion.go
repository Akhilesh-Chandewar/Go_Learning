package main

import "fmt"

func main() {
	fmt.Println("Recursion Example")
	n := 5
	fmt.Printf("Factorial of %d is %d\n", n, factorial(n))

	n = 2
	fmt.Printf("Sum of digits of %d is %d\n", n, sumOfDigits(n))

	n = 10
	fmt.Printf("Fibonacci of %d is %d\n", n, fibonacci(n))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sumOfDigits(n int) int{
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
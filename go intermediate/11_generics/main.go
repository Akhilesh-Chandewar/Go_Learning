package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s Stack[T]) printElements() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Stack elements:")
	for _, element := range s.elements {
		fmt.Println(element)
	}
}

func main() {
	a, b := 5, 10
	fmt.Println("Before swap:", a, b)
	a, b = swap(a, b)
	fmt.Println("After swap:", a, b)

	c, d := "hello", "world"
	fmt.Println("Before swap:", c, d)
	c, d = swap(c, d)
	fmt.Println("After swap:", c, d)

	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	intStack.printElements()
	fmt.Println("Stack size:", intStack.Size())
	intStack.Pop()
	fmt.Println("Stack size after pop:", intStack.Size())
	intStack.printElements()
	intStack.Pop()
}
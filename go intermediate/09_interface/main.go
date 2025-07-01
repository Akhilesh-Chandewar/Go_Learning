package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Rectangle1 struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle1) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle1) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g Geometry){
	fmt.Printf("Area: %.2f\n", g.area())
	fmt.Printf("Perimeter: %.2f\n", g.perimeter())
}

func main() {
	rectangle := Rectangle{width: 10, height: 5}
	circle := Circle{radius: 5}

	fmt.Println("Rectangle:")
	measure(rectangle)
	fmt.Println("Circle:")
	measure(circle)

	rectangle1 := Rectangle1{width: 10, height: 5}
	fmt.Println("Rectangle1:")
	measure(rectangle1) // This will not compile because Rectangle1 does not implement Geometry

	myPrinter("Hello", 42, 3.14, true)

	printType(42)
	printType("Hello")
	printType(3.14)
	printType(true)
	printType([]int{1, 2, 3})
	printType(struct{ name string }{"Go"})
	printType(func() { fmt.Println("Hello from function") })
	printType(nil)
	printType([]any{1, "two", 3.0, true})
	printType(map[string]int{"one": 1, "two": 2})
}

func myPrinter(i ...any){
	for _, v := range i {
		fmt.Println(v)
	} 
}

func printType(i any){
	fmt.Printf("%T\n", i)
}
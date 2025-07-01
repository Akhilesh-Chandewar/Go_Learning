package main

import "fmt"

type shape struct { Rectangle}

type Rectangle struct {
    length, width float64
}

func (r Rectangle) area() float64            { return r.length * r.width }
func (r *Rectangle) setLength(l float64)     { r.length = l }
func (r *Rectangle) setWidth(w float64)      { r.width = w }
func (r *Rectangle) scale(factor float64)    { r.length *= factor; r.width *= factor }

type MyInt int

func (m MyInt) isPositive() bool {
	return m > 0
}

func (MyInt) wellcomeMessage() string {
	return "Welcome to MyInt type!"
}

func main() {
    rect0 := Rectangle{length: 10, width: 5}
    fmt.Printf("Area of rectangle rect0: %.0f\n", rect0.area())

    var rect1 Rectangle
    rect1.setLength(20)
    rect1.setWidth(10)
    fmt.Printf("Area of rectangle rect1: %.0f\n", rect1.area())

    rect1.scale(2)
    fmt.Printf("Area of scaled rectangle rect1: %.0f\n", rect1.area())

	num := MyInt(5)
	if num.isPositive() {
		fmt.Println("num is positive")
	} else {
		fmt.Println("num is not positive")
	}
	num = MyInt(-3)
	if num.isPositive() {
		fmt.Println("num is positive")
	} else {
		fmt.Println("num is not positive")
	}

	fmt.Println(MyInt(0).wellcomeMessage())

	s := shape{Rectangle{length: 15, width: 10}}
	fmt.Printf("Area of shape: %.0f\n", s.area())
}

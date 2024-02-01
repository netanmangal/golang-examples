package main

import (
	"fmt"
	"math"
)

type IShapes interface {
	area() float64
}

type Rect struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	rect := Rect{height: 5.5, width: 4.0}
	circle := Circle{radius: 7}

	fmt.Println("rect is : ", rect, " - and circle is ", circle)

	fmt.Println("area of rect is : ", rect.area())
	fmt.Println("area of circle is : ", circle.area())

	fmt.Println("\nAfter interfaces : ")

	// shapes interface
	shapes := []IShapes{rect, circle}
	fmt.Println("shapes are : ", shapes)
	fmt.Println("After casting structs into interface, we can't access internal properties of struct. We can only use the methods and attributes defined in the interface.")

	for _, v := range shapes {
		fmt.Println("Inside FOR loop, area of the shape is : ", v.area())
	}
}


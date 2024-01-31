package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

type Triangle struct {
	v1 *Point
	v2 *Point
	v3 *Point
}

func (tri *Triangle) calcArea() float64 {
	v1, v2, v3 := tri.v1, tri.v2, tri.v3
	return math.Abs(float64(v1.x*(v2.y-v3.y)+v2.x*(v3.y-v1.y)+v3.x*(v1.y-v2.y))) / 2.0
}

func main() {
	tri := Triangle{v1: &Point{10, 20}, v2: &Point{20, 30}, v3: &Point{30, 80}}
	area := tri.calcArea()

	fmt.Println("Area of triangle with vertices : ", tri.v1, " , ", tri.v2, " , ", tri.v3, " is : ", area)
}


package main

import "fmt"

type Rectange struct {
	a, b int
}

func (r Rectange) Area() int {
	return r.a * r.b
}

func main() {
	var rect Rectange = Rectange{5, 6}
	fmt.Printf("Area of reactange with length - %d and breadth %d - is - %d", rect.a, rect.b, rect.Area())
}

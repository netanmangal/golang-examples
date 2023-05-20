package main

import "fmt"

func addAndReturn(a int, b int) float64 {
	return float64(a + b)
}

func main() {
	var a = 4
	var b int = 5

	c := addAndReturn(a, b)

	fmt.Printf("%d added to %d is %f, %v - which is stored as %T", a, b, c, c, c)
}

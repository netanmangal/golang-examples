package main

import "fmt"

func main() {
	a := 5
	fmt.Println("Value of a : ", a, " and addr of a : ", &a)

	a = 6
	fmt.Println("Value of a : ", a, " and addr of a : ", &a)

	aptr := &a
	*aptr = 100
	fmt.Println("Value of a : ", a, " and addr of a (aptr) : ", aptr)
}


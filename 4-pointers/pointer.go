package main

import "fmt"

func increase(i *int, by int) {
	*i = *i + by
}

func main() {
	var i int = 0

	increase(&i, 2)
	increase(&i, 3)

	fmt.Printf("Value of i is %d", i)
}

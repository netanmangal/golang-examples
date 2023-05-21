package main

import "fmt"

func main() {
	// declaration of map
	// // var m map[int8]string
	// this creates a nil map, this can only be read from and not be written to.

	m := make(map[int]string)

	m[0] = "User1"
	m[1] = "User2"
	m[2] = "User3"

	fmt.Println("map: ", m)
}

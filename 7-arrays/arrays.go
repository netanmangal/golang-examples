package main

import "fmt"

func main() {
	var arr [3]int

	arr[0], arr[1] = 2, 3
	arr[2] = 4

	fmt.Printf("Length of array is - %d \n\n", len(arr))

	for i, v := range arr {
		fmt.Printf("Value at index %d is %v \n", i, v)
	}

}

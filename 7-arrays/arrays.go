package main

import "fmt"

func main() {
	var arr [3]int
	fmt.Println("Empty array : ", arr)

	arr[1], arr[2] = 3, 4
	fmt.Println("Assigning values : ", arr)
	fmt.Println("Fetch value of arr[2] : ", arr[2])
	fmt.Println("Length of array : ", len(arr))

	fmt.Println("\nNext Array")

	barr := [5]int{11, 22, 33, 44, 55}
	fmt.Println("Decl array barr : ", barr)

	for i, v := range barr {
		fmt.Println("From for loop, Value at index : ", i, " is : ", v)
	}
}


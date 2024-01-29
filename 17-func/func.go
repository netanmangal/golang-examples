package main

import "fmt"

func findSum(inputs ...int) int {
	sum := 0
	fmt.Println("Received : ", len(inputs), " elements to sum up")

	for _, v := range inputs {
		sum += v
	}

	return sum
}

func recursiveFibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return (recursiveFibonacci(n-1) + recursiveFibonacci(n-2))
	}
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of slice1 : ", slice1, " : is : ", findSum(slice1...))

	n := 10
	fmt.Println("\n\nFibonacci number at ", n, " is : ", recursiveFibonacci(n))
}


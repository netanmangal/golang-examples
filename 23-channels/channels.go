package main

import "fmt"

func main() {
	n := 100
	ch := make(chan int)

	go func(nos int) {
		sum := 0
		for i := 1; i <= nos; i++ {
			sum += i
		}
		ch <- sum
	}(n)

	sum := <-ch
	fmt.Println("Sum of first ", n, " natural numbers is : ", sum)
}

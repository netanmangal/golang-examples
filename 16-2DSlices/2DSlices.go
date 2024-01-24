package main

import "fmt"

func main() {
	twoDSlice := make([][]int, 4)

	for i, _ := range twoDSlice {
		twoDSlice[i] = make([]int, i+3)

		for j, _ := range twoDSlice[i] {
			twoDSlice[i][j] = i + j + 1
		}
	}

	fmt.Println(twoDSlice)
}


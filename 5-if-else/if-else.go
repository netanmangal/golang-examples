package main

import "fmt"

func main() {
	passingScores := 33

	var scored int = 30

	if scored > passingScores {
		fmt.Printf("You passed, Good job!!")
	} else if scored == passingScores {
		fmt.Printf("You are barely passing. Work hard!!")
	} else {
		fmt.Printf("You failed. Work very hard!!")
	}
}

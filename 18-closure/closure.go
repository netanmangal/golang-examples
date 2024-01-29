package main

import "fmt"

func initializeSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println("Sequence 01 : ")
	incrementSeq01 := initializeSeq()
	fmt.Println(incrementSeq01())
	fmt.Println(incrementSeq01())
	fmt.Println(incrementSeq01())

	fmt.Println("\nSequence 02 : ")

	incrementSeq02 := initializeSeq()
	fmt.Println(incrementSeq02())
	fmt.Println(incrementSeq02())
	fmt.Println(incrementSeq02())

}


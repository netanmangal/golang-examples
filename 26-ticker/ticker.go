package main

import (
	"fmt"
	"time"
)

func timeTicker() {
	ticker := time.NewTicker(time.Second)

	for {
		ch := <- ticker.C
		fmt.Println("Printing it after - ", ch, " seconds.")

		// if don't want to read the value in `ch`, we can do something like
		// <- ticker.C
		// just blocking the loop till the channel ticker.C is going to have value
	}
}

func main() {
	timeTicker()
}
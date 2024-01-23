package main

import (
	"fmt"
	"time"
)

func printString(s string) {
	fmt.Println(s)
}

func main() {
	today := time.Now()
	switch w:= today.Weekday(); w {
	case time.Saturday:
		printString("Hurray, it's Saturday")
	case time.Sunday:
		printString("OH, it's Sunday, let's pack the bag")
	default:
		printString(fmt.Sprintf("Keep working... it's not the weekend. it's %s", w))
	}

	switch h := today.Hour(); {
	case h < 12:
		printString("It's before noon !")
	case h == 12:
		printString("It's noon !")
	case h > 12:
		printString("It's after noon !")
	}

	fmt.Println("\n --- It's ", today, " ---")
}


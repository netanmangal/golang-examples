package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Please input - ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("Your input is - ", scanner.Text())
}
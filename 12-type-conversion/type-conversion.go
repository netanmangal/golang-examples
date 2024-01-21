package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the year you were born : ")
	scanner.Scan()
	user_age, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	fmt.Println("You entered : ", user_age)

	fmt.Println("Your age is : ", 2024-user_age)
}

package main

import "fmt"

func main() {
	var i interface{}

	i = 4

	switch i.(type) {
	case string:
		fmt.Printf("It's string - the value is %v", i.(string))

	case int:
		fmt.Printf("It's int - the value is %v", i.(int))

	case float64:
		fmt.Printf("It's float64 - the value is %v", i.(float64))

	default:
		fmt.Printf("It's of custom type - %T - the value is %v", i, i)
	}
}

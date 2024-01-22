package main

import "fmt"

func main() {
	fmt.Println("Most basic FOR loop, with only one condition.")

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Classic FOR loop, with initialization | condition | after-step.")
	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	fmt.Println("Using continue and break.")

	for k := 0; ; k++ {
		if k%2 == 0 {
			fmt.Println(k, " is even number")
		} else {
			if k > 10 {
				break
			} else {
				continue
			}
		}
	}
}


package main

import (
	"fmt"
	"slices"
)

func main() {
	var s1 []string
	fmt.Println("Empty slice1 : ", s1)
	fmt.Println("slice1 : len(s1) == 0 : ", len(s1) == 0)
	fmt.Println("slice1 : s1 == nil    : ", s1 == nil)

	s1 = make([]string, 5)
	fmt.Println("After allocation cap of 5, len(s1) : ", len(s1))
	fmt.Println("Empty slice s1 : ", s1)

	s1[0], s1[1], s1[2], s1[3], s1[4] = "xmas", "yawk", "apple", 
"boy", "cat"
	fmt.Println("After set, s1 : ", s1)

	s1 = append(s1, "hen")
	fmt.Println("Append to slice, s1 : ", s1, " with new len : ", 
len(s1))

	// slice operations
	fmt.Println()
	fmt.Println("s1      : ", s1)
	fmt.Println("s1[1:3] : ", s1[1:3])
	fmt.Println("s1[ :4] : ", s1[ :4])
	fmt.Println("s1[3: ] : ", s1[3: ])

	s2 := make([]string, len(s1))
	copy (s2, s1)
	fmt.Println("s2 copied from s1, s2 : ", s2)	
	
	if slices.Equal(s1, s2) {
		fmt.Println("slices.Equal(s1, s2) : true")
	} else {
		fmt.Println("slices.Equal(s1, s2) : false")
	}
}


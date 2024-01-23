package main

import "fmt"

func main() {
	arr := [3][5]int{{1, 2, 3, 4, 5}, {3, 4, 5, 6, 7}, {6, 7, 8, 9, 0}}

	fmt.Println("arr : ", arr)
	fmt.Println("# of rows : ", len(arr))
	fmt.Println("# of columns : ", len(arr[0]))

	fmt.Println("\n Next array")

	var barr [2][4]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			barr[i][j] = i + j
		}
	}

	fmt.Println("barr : ", barr)
	fmt.Println("# of rows : ", len(barr))
	fmt.Println("# of columns : ", len(barr[0]))
}


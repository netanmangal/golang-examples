package main

import "fmt"

func main() {
	m1 := map[int]string{}
	m1[0], m1[1] = "hello", "world"

	fmt.Println("map m1 : ", m1)
	fmt.Printf("type of map m1 : %T \n", m1)
	fmt.Println("len of map m1 : ", len(m1))

	m2 := map[int]string{3: "apple", 4: "banana", 5: "carrot", 6: "orange"}
	fmt.Println("map m2 : ", m2)

	delete(m2, 5)
	fmt.Println("deleting key[5] from m2 : ", m2)

	m2[7] = ""
	fmt.Println("map m2 : ", m2)

	_, present := m2[7]
	fmt.Println("Optional return value 2 from m2[7] if key is present : ", present)
}


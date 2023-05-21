package main

import "fmt"

type Gender int8

const (
	Male   Gender = 0
	Female Gender = 1
	Other  Gender = 2
)

type Person struct {
	name   string
	age    int
	gender Gender
}

func main() {
	// we can initialize some elements while declaration itself
	persons := []Person{
		{"rahul", 23, Male},
	}

	persons = append(persons, Person{"raghav", 34, Female})

	fmt.Printf("There are %d persons \n\n", len(persons))

	for i, v := range persons {
		fmt.Printf("Person's details at index %d - Name: %s Age: %d Gender: %v \n\n", i, v.name, v.age, v.gender)
	}
}

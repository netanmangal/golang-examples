package main

import (
	"fmt"

	"github.com/netanmangal/golang-examples/25-oops/composition"
	"github.com/netanmangal/golang-examples/25-oops/polymorphism"
	"github.com/netanmangal/golang-examples/25-oops/types"
)

func main() {
	fmt.Println("\n\n---- Encapsulation ----")

	person := types.Person{}
	person.SetPersonAttributes("Netan", "Mangal", "12-12-12", "M", []string{"trekking", "playing guitar"})
	person.PrintAllDetails()

	fmt.Println("\n\n---- Polymorphism ----")
	
	// here, on the left hand side, the type of variables is Shapes,
	// but we're assigning them the type of Circle and Rectangle, 
	// because the structs have implemented the type Shapes
	// which essentially means to implement receiver functions of the interface
	var circle polymorphism.Shapes = &polymorphism.Circle{Radius: 5.0}
	var rect polymorphism.Shapes = &polymorphism.Rectangle{Width: 10.0, Height: 7.0}

	fmt.Println("Area of circle : ", circle.Area())
	fmt.Println("Area of rect : ", rect.Area())

	fmt.Println("\n\n---- Composition ----")
	var car *composition.Car = composition.NewCar(1750, 22, "new engine", "tata wheel")
	fmt.Println(" - Accessing the properties and function -")
	fmt.Println("Car's engine details :: hp : ", car.GetHP(), " name : ", car.GetName())
	fmt.Println("Car's wheels details :: dimensions : ", car.Wheel.GetDimensions(), " name : ", car.Wheel.GetName())


	fmt.Println("\n\n---- Comparable ----")
	person2 := types.Person{}
	person2.SetPersonAttributes("Netan", "Mangal", "12-12-12", "M", []string{"trekking", "playing guitar"})
	
	fmt.Println("Person 1 : ", person)
	fmt.Println("Person 2 : ", person2)

	// we can only compare with == operation if struct has primitive datatypes
	// if person == person2 {
	// 	fmt.Println("Two structs are equal.")
	// } else {
	// 	fmt.Println("Structs are not equal.")
	// }

	fmt.Println("Person 1 equals Person 2 : ", person.Equals(&person2))
}
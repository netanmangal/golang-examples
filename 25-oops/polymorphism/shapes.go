package polymorphism

import "math"

type Shapes interface{
	Area() float32
}

type Rectangle struct {
	Height float32
	Width float32
}

type Circle struct {
	Radius float32
}

func (r *Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (c *Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

package composition

type engine struct {
	horsepower int
	name string
}

func (e *engine) GetHP() int {
	return e.horsepower
}

func (e *engine) GetName() string {
	return e.name
}

type wheel struct {
	wheelDimension int
	name string
}

func (w *wheel) GetDimensions() int {
	return w.wheelDimension
}

func (w *wheel) GetName() string {
	return w.name
}

// this is composition - `has a` relation-ship
// we can use composition in 2 ways - 
// 1) without naming the attribute
// 2) by naming the attribute of embedded struct
type Car struct {
	engine
	Wheel wheel
}

func NewCar(hp, wd int, egName, whName string) *Car {
	return &Car{
		engine: engine{
			horsepower: hp,
			name: egName,
		},
		Wheel: wheel{
			wheelDimension: wd,
			name: whName,
		},
	}
}

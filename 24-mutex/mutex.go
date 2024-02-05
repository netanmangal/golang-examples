package main

import (
	"sync"
)

type Structure struct {
	lock *sync.Mutex
	value map[string]int
}

func NewStructure() *Structure {
	return &Structure{
		lock: &sync.Mutex{},
		value: map[string]int{ "x": 5 },
	}
}

func (s *Structure) updateValue () {
	s.lock.Lock()
	defer s.lock.Unlock()
	
	s.value["x"] = 1
}

func main() {
	newStruct := NewStructure()

	for i := 0; i < 100; i++ {
		go func () {
			newStruct.updateValue()
		}()
	}
}
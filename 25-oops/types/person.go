package types

import "fmt"

// this is encapsulation
// attributes with first letter capitalized are Public attributes
// and can be accessed outside the package
// attributes with first letter small-cased are private
// and can be accessed within the package only
// this applies to "first-letter" of receiver functions as well

type Person struct {
	FirstName string
	LastName string
	dob string
	gender string
	hobbies []string
}

// pointer based receiver function
func (p *Person) SetPersonAttributes (firstname, lastname, dob, gender string, hobbies []string) {
	p.FirstName = firstname
	p.LastName = lastname
	p.dob = dob
	p.gender = gender
	p.hobbies = hobbies
}

func (p *Person) GetDoB() string {
	return p.dob
}

func (p *Person) GetGender() string {
	return p.gender
}

func (p *Person) GetHobbies() []string {
	return p.hobbies
}

func (p *Person) PrintAllDetails() {
	fmt.Println(p)
}

func (p1 *Person) Equals(p2 *Person) bool {
	if p1.FirstName != p2.FirstName ||
		p1.LastName != p2.LastName ||
		p1.gender != p2.gender ||
		p1.dob != p2.dob {
		return false
	}

	if len(p1.hobbies) != len(p2.hobbies) {
		return false
	}

	for i, v := range p1.hobbies {
		if p2.hobbies[i] != v {
			return false
		}
	}

	return true
}

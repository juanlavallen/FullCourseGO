package models

type Person struct {
	// Name string // property public
	name string // property private
	age  int    // property private
}

// Constructor
func (p *Person) Constructor(name string, age int) {
	p.name = name
	p.age = age
}

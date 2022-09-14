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

// GET
func (p *Person) GetName() string {
	return p.name
}

// SET
func (p *Person) SetName(name string) {
	p.name = name
}

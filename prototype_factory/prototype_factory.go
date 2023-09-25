package copy_through_serialization

import (
	"bytes"
	"encoding/gob"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

var mainPerson = Person{
	"John",
	&Address{"123" +
		"London Rd", "London", "UK"},
	[]string{"Chris", "Matt"},
}

var auxPerson = Person{
	"Jane",
	&Address{"123" +
		"London Rd", "London", "UK"},
	[]string{"Chris", "Matt"},
}

func newPerson(proto *Person, name string) *Person {
	result := proto.DeepCopy()
	result.Name = name
	return result
}

func NewMainPerson(name string) *Person {
	return newPerson(&mainPerson, name)
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

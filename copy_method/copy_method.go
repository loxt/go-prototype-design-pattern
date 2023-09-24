package deep_copying

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country}
}

func main() {
	john := Person{"John", &Address{"123 London Rd", "London", "UK"}, []string{"Chris", "Matt"}}
	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "124 London Rd"
	jane.Friends = append(jane.Friends, "Angela")
}

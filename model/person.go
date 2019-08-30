package modeldb

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) GetFirstName() string {
	return p.FirstName
}

func (p *Person) ChangeFirstName(firstname string) {
	p.FirstName = firstname
}

package modeldb

type Person struct {
	FirstName string `json:"first_name" bson:"first_name" gorm:"column:firstname" validate:"email,required"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Married   bool   `json:"married"`
}

func (p Person) GetFirstName() string {
	return p.FirstName
}

func (p *Person) ChangeFirstName(firstname string) {
	p.FirstName = firstname
}

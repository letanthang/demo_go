package db1

type Student struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	ClassName   string `json:"course_name"`
	AcademyName string `json:"acedemy_name"`
}

func (s Student) GetName() string {
	return s.FirstName + " " + s.LastName
}

func (s *Student) SetAge(age int) {
	s.Age = age
}

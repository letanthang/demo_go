package main

import (
	"fmt"

	db2 "github.com/letanthang/demo_go/model"
)

func main() {
	aStudent := db2.Student{FirstName: "Sang", LastName: "Nguyen", Age: 100}

	fmt.Println(aStudent.GetName())
	// SetAge - 1000
	aStudent.SetAge(1000)

	var p *db2.Student
	p = &aStudent
	p.Age = 20000
	fmt.Println(aStudent.Age)
}

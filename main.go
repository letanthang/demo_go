package main

import (
	"fmt"

	db "github.com/letanthang/demo_go/model"
)

func main() {
	bs := []int{10, 100, 2000, 2000, 100, 2000, 10}

	var aMap map[int]int
	aMap = map[int]int{}

	for _, v := range bs {
		aMap[v] = v
	}

	var uniqueSlice []int
	for v := range aMap {
		uniqueSlice = append(uniqueSlice, v)
	}
	fmt.Println(bs, aMap, uniqueSlice)

	aPerson := db.Person{FirstName: "Thang", LastName: "Le"}
	aPerson.ChangeFirstName("Tam")
	firstName := aPerson.GetFirstName()
	fmt.Println(firstName)
}

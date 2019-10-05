package main

import "fmt"

func main() {

	// aPerson := Person{Name: "Thai", Age: 90}
	aBoy := Boy{Name: "Tam", Age: 100}

	PrintPersonName(Person(aBoy))

}
func PrintPersonName(p Person) {
	fmt.Println(p.Name)
}

type Person struct {
	Name string
	Age  int `json:"age"`
}

type Boy struct {
	Name string
	Age  int `json:"abc"`
}

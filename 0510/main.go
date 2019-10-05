package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	ClassName   string `json:"course_name"`
	AcademyName string `json:"acedemy_name"`
}

func main() {

	//parse string to struct
	jsonString := "{\"first_name\":\"Tin\",\"last_name\":\"Tran\",\"email_address\":\"tintran@gmail.com\",\"age\":80,\"course_name\":\"golang0110\",\"acedemy_name\":\"Nordic Coder\"}"
	fmt.Println(jsonString)

	bs := []byte(jsonString)
	var aStudent Student
	json.Unmarshal(bs, &aStudent)

	//type casting
	//1. int - float
	aInt := 5
	var aFloat float64
	aFloat = float64(aInt)
	fmt.Println("aFloat=", aFloat)

	//2. struct -struct
	aBoy := Boy{Name: "Thai", Age: 90}
	PrintPersonName(Person(aBoy))

	//type assertion
	var i interface{}
	i = "62.0"
	var bFloat float64
	bFloat, _ = i.(float64)
	fmt.Println("bFloat=", bFloat)

}

type Person struct {
	Name string
	Age  int
}

type Boy struct {
	Name string
	Age  int
}

func PrintPersonName(p Person) {
	fmt.Println(p.Name)
}

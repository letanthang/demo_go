package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString := `{
		"first_name": "Van",
		"last_name": "Tran",
		"email_address": "hoangvan060@gmail.com",
		"age": 80,
		"course_name": "golang0110",
		"acedemy_name": "Nordic Coder"
	}`

	var aMap map[string]interface{}

	bs := []byte(jsonString)
	json.Unmarshal(bs, &aMap)
	// fmt.Println(aMap)

	var age interface{}
	age = aMap["age"]
	fmt.Println(age)

	var numAge float64
	// if numAge, ok := age.(float64); ok == false {
	// 	fmt.Println("cannot type assertion", numAge)
	// }
	numAge = age.(float64)
	fmt.Println(numAge)

}

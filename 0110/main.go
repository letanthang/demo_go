package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type MyInt int

type Student struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	ClassName   string `json:"class_name"`
	AcademyName string `json:"acedemy_name"`
}

func main() {
	fmt.Println("Hello World Golang 01/10/2019")
	// int
	var aInt int
	aInt = 10
	fmt.Println("aInt=", aInt)
	bInt := 11
	fmt.Println("bInt=", bInt)
	bInt = 12
	fmt.Println("bInt=", bInt)

	// string
	var aString string
	aString = "hello"
	bString := "1000"
	fmt.Println(aString)
	// string concat : way 1
	aString = aString + bString
	fmt.Println(aString)

	fmt.Printf("%s\n", aString)
	fmt.Printf("%d\n", aInt)
	// string concat : way 2
	aString = fmt.Sprintf("hehe %s %d\n", aString, aInt)
	fmt.Println(aString)

	res, err := getAnInterger(bString)
	fmt.Println(res, err)

	aSlice := []int{1, 5, 6, 8}
	fmt.Println(aSlice)
	fmt.Printf("%v", aSlice)
	// slice iterate
	for _, v := range aSlice {
		fmt.Println(v)
	}

	aMap := map[string]interface{}{"name": "Phat", "email": "Phat@gmail.com"}
	aMap["age"] = 100
	fmt.Printf("%v", aMap)

	// map iterate
	for k, v := range aMap {
		fmt.Println("key", k, "has value", v)
	}

	aStudent := Student{
		FirstName:   "Tin",
		LastName:    "Tran",
		Email:       "tintran@gmail.com",
		Age:         80,
		ClassName:   "golang0110",
		AcademyName: "Nordic Coder",
	}
	bs, _ := json.Marshal(aStudent)
	dString := string(bs)
	fmt.Printf("%v+ \n %s", aStudent, dString)
}
func getAnInterger(a string) (int, error) {
	res, _ := strconv.Atoi(a)
	return res, errors.New("bad request")
}

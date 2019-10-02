package main

import (
	"errors"
	"fmt"
	"strconv"
)

type MyInt int
type Student struct {
	Name      string `json:"name"`
	Email     string
	Age       int
	ClassName string `json:"class_name"`
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
	aString = aString + bString
	fmt.Println(aString)

	fmt.Printf("%s\n", aString)
	fmt.Printf("%d\n", aInt)

	aString = fmt.Sprintf("hehe %s %d\n", aString, aInt)
	// aString = aString + " " + strconv.Itoa(aInt)
	fmt.Println(aString)

	res, err := getAnInterger(bString)
	fmt.Println(res, err)

	aSlice := []int{1, 5, 6, 8}
	fmt.Println(aSlice)
	fmt.Printf("%v", aSlice)

	for _, v := range aSlice {
		fmt.Println(v)
	}

	aMap := map[string]interface{}{"name": "Phat", "email": "Phat@gmail.com"}

	aMap["age"] = 100
	fmt.Printf("%v", aMap)

	aStudent := Student{
		Name:  "Luc",
		Email: "luc@gmail.com",
		Age:   80,
	}
	fmt.Printf("%v+", aStudent)
}
func getAnInterger(a string) (int, error) {
	res, _ := strconv.Atoi(a)
	return res, errors.New("bad request")
}

package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	aInt := 5
	aType := reflect.TypeOf(aInt)
	fmt.Printf("type %s | kind %s \n", aType.Name(), aType.Kind())
	// type int , kind int

	aStruct := struct {
		Name string
		Age  int
	}{Name: "Thang", Age: 35}

	aType = reflect.TypeOf(aStruct)
	fmt.Printf("type %s | kind %s \n", aType.Name(), aType.Kind())
	// type nil , kind struct

	// aPerson := Person{Name: "T", Age: 35}

}

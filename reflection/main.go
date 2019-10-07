package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Foo struct {
	A int    `json:"a" gorm:"colum:a" tag1:"hehe" tag2:"hihi"`
	B string `json:"b" gorm:"colum:b"`
}

func main() {
	// struct
	// type struct
	// slice
	// pointer
	// map
	// get kind and type
	// case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice : get element: Elem()
	// case reflect.Struct : get field : NumField(), Field(int)
	aStruct := struct {
		A int
		B string
	}{1, "hehe"}

	bStruct := Foo{A: 2, B: "hihi"}
	cSlice := []int{1, 2, 4}
	dPointer := &cSlice
	eMap := map[string]interface{}{"name": "Thang", "age": 1}

	aType := reflect.TypeOf(aStruct)
	bType := reflect.TypeOf(bStruct)
	cType := reflect.TypeOf(cSlice)
	dType := reflect.TypeOf(dPointer)
	eType := reflect.TypeOf(eMap)

	fmt.Println("bType name=", aType.Name(), "bType Kind=", aType.Kind())
	fmt.Println("bType name=", bType.Name(), "bType Kind=", bType.Kind(), "fields=", bType.Field(0))
	pointer := &bType
	fmt.Println("type of", bType, pointer, reflect.ValueOf(bType), reflect.ValueOf(pointer))

	b1 := reflect.New(bType)
	b1.Elem().Field(0).SetInt(3)
	b1.Elem().Field(1).SetString("kaka")
	b1Struct := b1.Elem().Interface().(Foo)
	fmt.Println(b1Struct.A, b1Struct.B)

	fmt.Println("------------------------------")
	inspect(aType, 0)
	inspect(bType, 0)
	inspect(cType, 0)
	inspect(dType, 0)
	inspect(eType, 0)

	fmt.Println("ba" + strings.Repeat("na", 5))
}

func inspect(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "- kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		inspect(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name)
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "Tag1 is", f.Tag.Get("tag1"), "Tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}

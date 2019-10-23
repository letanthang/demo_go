package main

import "errors"

func main() {

}

type Bar struct {
	ID   int
	Name string
	Err  error
}

func (b Bar) Error() string {
	return b.Err.Error()
}

func Foo() error {
	res := Bar{1, "Thang", errors.New("Some Error")}
	return res
}

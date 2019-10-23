package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	ID        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

func main() {
	url := "http://localhost:9090/api/v1/public/student"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var result []Student // var result []map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Students : %+v", result)
}

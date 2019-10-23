package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	url := "http://service-name/api/v1/public/student"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		panic(err)
	}

	fmt.Printf("students: %+v", result)

}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	model "github.com/letanthang/demo_go/model"

	"github.com/letanthang/demo_go/db"
)

func main() {

	people, err := GetPeopleNew()
	if err != nil {
		log.Panic(err)
		return
	}
	err = StorePeople(*people)
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println("************** success ****************")
	fmt.Printf("people: %+v", people)
}

func StorePeople(people []model.Person) error {
	sessionClone := db.SessionOrginal.Clone()
	defer sessionClone.Close()

	err := sessionClone.DB("go3008").C("people").Insert(people[0])
	// bulk := sessionClone.DB("go3008").C("people").Bulk()
	// bulk.Insert(*people...)
	return err

}

func GetPeople() (*[]model.Person, error) {
	url := "http://localhost:9090/v1/people"
	data := map[string]interface{}{"course": "golang"}
	bs, _ := json.Marshal(data)
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer(bs))
	req.Header.Add("Authorization", "key=AAAAUxSNhvw:APA91bGdIHmDV9ANR7zdhGbbIfaOSQgIq8ziI0djh0wH1nDGXfhRomrYx7zG_18vQZ1T5eqSVVlCBaL3hhIgHa0MK49eF103GlbIniL8BQUKt7tMDprqZZ_-UzA-frx7n8M9aj7Kgbfv")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var people []model.Person
	err = json.Unmarshal(bs, &people)
	if err != nil {
		return nil, err
	}
	return &people, nil
}

func GetPeopleNew() (*[]model.Person, error) {
	url := "http://localhost:9090/v1/people"
	data := map[string]interface{}{"course": "golang"}
	bs, _ := json.Marshal(data)
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer(bs))
	req.Header.Add("Authorization", "key=AAAAUxSNhvw:APA91bGdIHmDV9ANR7zdhGbbIfaOSQgIq8ziI0djh0wH1nDGXfhRomrYx7zG_18vQZ1T5eqSVVlCBaL3hhIgHa0MK49eF103GlbIniL8BQUKt7tMDprqZZ_-UzA-frx7n8M9aj7Kgbfv")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var people []model.Person
	err = decoder.Decode(&people)

	return &people, nil
}

func practise() {
	bs := []int{10, 100, 2000, 2000, 100, 2000, 10}

	var aMap map[int]int
	aMap = map[int]int{}

	for _, v := range bs {
		aMap[v] = v
	}

	var uniqueSlice []int
	for v := range aMap {
		uniqueSlice = append(uniqueSlice, v)
	}
	fmt.Println(bs, aMap, uniqueSlice)

	aPerson := model.Person{FirstName: "Thang", LastName: "Le"}
	aPerson.ChangeFirstName("Tam")
	firstName := aPerson.GetFirstName()
	fmt.Println(firstName)
}

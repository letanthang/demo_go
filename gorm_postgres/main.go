package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

type Student struct {
	ID        int
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Age       int
	ClassName string `json:"class_name" gorm:"column:class_name"`
	Email     string `json:"email" gorm:"column:email"`
}

func (Student) TableName() string {
	return "student"
}

func main() {
	DB = newDB()

	var students []Student
	DB.Find(&students)
	fmt.Printf("students :%+v", students)
}

func newDB() *gorm.DB {
	// db, err := gorm.Open("sqlite3", "test.db")
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "localhost", "postgres", "nc", "docker")
	fmt.Println(connectionString)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().Ping()
	db.LogMode(true)

	return db
}

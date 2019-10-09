package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 5)
	go DoTask2(ch)
	go DoTask3(ch)
	go DoTask1(ch)

	<-ch
	<-ch
	<-ch
	<-ch

}

func DoTask1(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Task1 done")
	ch <- 1
}

func DoTask2(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Task2 done")
	ch <- 1
}

func DoTask3(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Task3 done")
	ch <- 1
}

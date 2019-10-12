package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go DoTask1(ch)
	go DoTask2(ch)
	go DoTask3(ch)
	<-ch
	<-ch

}

func DoTask1(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Done1")
	ch <- 1
}

func DoTask2(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Done2")
	ch <- 1
}

func DoTask3(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Done3")
	ch <- 1
}

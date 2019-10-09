package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go DoTask2(&wg)

	wg.Add(1)
	go DoTask3(&wg)

	DoTask1()
	wg.Wait()
}

func DoTask1() {
	time.Sleep(2 * time.Second)
	fmt.Println("Task1 done")
}

func DoTask2(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Task2 done")
	wg.Done()
}

func DoTask3(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("Task3 done")
}

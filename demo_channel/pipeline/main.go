package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// ch := make(chan int)

	// go func() {
	// 	ch <- 100
	// }()

	// var aInt int
	// aInt = <-ch
	// fmt.Println("read value:", aInt)

	ch := asChan(1, 2, 3, 4, 5, 6)
	c1 := AddOne(ch)
	c2 := DoubleValue(c1)

	// 4 6 8 10 12 14
	fmt.Println("channel value :")
	for v := range c2 {
		fmt.Print(v, " ")
	}

}

func AddOne(c chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out) // đóng channel
		for v := range c {
			out <- v + 1
		}
	}()
	return out
}

func DoubleValue(c chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out) // đóng channel

		for v := range c {
			out <- v * 2
		}
	}()
	return out
}
func asChan(vs ...int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, v := range vs {
			out <- v
			time.Sleep(100 * time.Millisecond)
		}
	}()

	return out
}

// var wg sync.WaitGroup
// wg.Add(3)
// go DoTask2(&wg)
// go DoTask3(&wg)
// go DoTask1(&wg)
// wg.Wait()
func DoTask1(wg *sync.WaitGroup) {
	defer wg.Done()
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

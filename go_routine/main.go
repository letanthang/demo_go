package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(1)
	go DoSomething2(&waitgroup)

	waitgroup.Add(1)
	go DoSomething3(&waitgroup)

	waitgroup.Add(1)
	DoSomething1(&waitgroup)
	waitgroup.Wait()
}

func DoSomething1(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Done1")
	wg.Done()
}
func DoSomething2(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Done2")
	wg.Done()
}

func DoSomething3(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Done3")
	wg.Done()
}

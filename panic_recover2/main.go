package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		defer close(ch)
		ch <- 1
		time.Sleep(1 * time.Second)
		ch <- 2
		time.Sleep(1 * time.Second)
		ch <- 0
		time.Sleep(1 * time.Second)
		ch <- 10
		time.Sleep(1 * time.Second)

	}()

	for i := range ch {
		safelyDo(i)
	}

}

func safelyDo(work int) {
	// call recover() in defer
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	dangerous(work)
}

func dangerous(work int) {

	res := 100 / work
	fmt.Print(res, " ")
}

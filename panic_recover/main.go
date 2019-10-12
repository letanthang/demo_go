package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	ch := asChan(1, 2, 4, 0, 5, 6)

	for v := range ch {
		// fmt.Print(v, " ")
		safelyDo(v)
	}
}

func asChan(vs ...int) chan int {
	out := make(chan int, 2)
	go func() {
		defer close(out)
		for _, v := range vs {
			time.Sleep(20 * time.Millisecond)
			out <- v
		}
	}()

	return out
}

func safelyDo(work int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	do(work)
}

func do(work int) {
	res := 100 / work
	fmt.Print(res, " ")
}

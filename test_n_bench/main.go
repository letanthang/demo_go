package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := AsChan(1, 2, 3, 4, 5)
	for v := range ch {
		fmt.Println(v)
	}

	ch1 := AsChan(1, 2, 3)
	ch2 := AsChan(4, 5, 6)
	ch3 := AsChan(7, 8, 9)

	mCh := MergeChannel(ch1, ch2, ch3)
	fmt.Println("merged channel:")
	for v := range mCh {
		fmt.Print(v, " ")
	}
}

func MergeChannel(chans ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		var wg sync.WaitGroup
		for _, ch := range chans {
			wg.Add(1)
			go func(c <-chan int) {
				defer wg.Done()
				for v := range c {
					out <- v
				}
			}(ch)
		}
		wg.Wait()
	}()
	return out
}

func AsChan(nums ...int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for _, v := range nums {
			c <- v
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
		}
	}()

	return c
}

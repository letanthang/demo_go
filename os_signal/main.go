package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	forever := make(chan os.Signal)
	signal.Notify(forever, os.Interrupt)
	go func() {
		i := 0
		for {
			i++
			fmt.Println("Hello world", i)
			time.Sleep(800 * time.Millisecond)
		}
	}()

	<-forever
	fmt.Println("Thanks for learning Golang")
}

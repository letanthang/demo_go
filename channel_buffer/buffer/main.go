package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "some data"
	ch <- "other data"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

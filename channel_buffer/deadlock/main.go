package main

import "fmt"

func main() {
	ch := make(chan string)

	ch <- "some data"

	fmt.Println(<-ch)
}

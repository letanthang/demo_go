package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "https://vnexpress.net/"
	}()

	resultChannel := startWorker(ch)

	urlMap := map[string]bool{}

	for v := range resultChannel {
		fmt.Println("Receive 1 result", v)
		for _, url := range v.Urls {
			if urlMap[url] != true {
				ch <- url
				urlMap[url] = true
			}

		}
	}
}

func startWorker(ch chan string) chan ParseResult {
	out := make(chan ParseResult)
	go func() {
		guard := make(chan int, 4)
		for url := range ch {
			guard <- 1
			go func(url string) {
				pr := parse(url)
				out <- pr
				<-guard
			}(url)
		}
	}()

	return out
}

type ParseResult struct {
	Title  string
	Author string
	Urls   []string
}

func parse(url string) ParseResult {
	return ParseResult{}
}

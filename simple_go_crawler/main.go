package main

import "fmt"

func main() {

	ch := make(chan string)

	go func() {
		ch <- "vnexpress.net/home"
	}()

	resultChannel := startWorker(ch)

	for result := range resultChannel {
		fmt.Println("Fin 1 result", result)

		for _, url := range result.Urls {
			go func(url string) {
				ch <- url
			}(url)
		}
	}

}

type CrawlResult struct {
	Content string
	Title   string
	Author  string
	Urls    []string
}

func startWorker(ch chan string) chan CrawlResult {
	out := make(chan CrawlResult)
	go func() {
		for url := range ch {
			go func(url string) {
				result := parseUrl(url)
				out <- result
			}(url)
		}
	}()

	return out
}

func parseUrl(url string) CrawlResult {
	return CrawlResult{
		Content: "Indo thua Viet Nam 3-1",
		Title:   " Indo thua",
		Author:  "Tung Son",
		Urls:    []string{"vnexpress.net/1", "vnexpress.net/2", "vnexpress.net/3", "vnexpress.net/4"},
	}
}

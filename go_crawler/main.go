package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "https://vnexpress.net/the-gioi/dong-co-may-bay-boc-chay-khi-sap-cat-canh-3999385.html?vn_source=Home&vn_campaign=TopNews&vn_medium=Item-1&vn_term=Desktop&vn_thumb=0"
	}()

	resultChannel := startWorker(ch)

	for result := range resultChannel {
		fmt.Println("Find on result!", result)
		// write workload/urls to ch
		go func() {
			for _, url := range result.Urls {
				ch <- url
			}
		}()
	}
}

type Result struct {
	Title       string
	Description string
	Content     string
	Urls        []string
}

func startWorker(ch chan string) chan Result {
	out := make(chan Result)
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

func parseUrl(url string) Result {
	return Result{
		Title:       "Viet Nam thang Indo 3-1",
		Description: "lasjd;fals;dfj",
		Content:     "HÀN QUỐCĐộng cơ máy bay hãng Asiana Airlines bốc cháy trên đường băng sân bay Incheon ở Seoul khi đang tiếp nhiên liệu để chuẩn bị cất cánh. Hành khách đang chờ lên máy bay tại sân bay quốc tế Incheon hôm 18/10 hoảng sợ khi chứng kiến ngọn lửa bùng lên từ một động cơ của chiếc Airbus A380. Động cơ bốc cháy trong lúc máy bay đang tiếp nhiên liệu để khởi hành từ Seoul tới Los Angeles, Mỹ với 495 hành khách.Không có ai bị thương trong sự cố và toàn bộ hành khách đã được sơ tán khỏi khu vực cổng ra máy bay.",
		Urls:        []string{"https://vnexpress.net/the-gioi/1.html", "https://vnexpress.net/the-gioi/2.html", "https://vnexpress.net/the-gioi/3.html"},
	}
}

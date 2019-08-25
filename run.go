package crawler

import (
	"log"
	"sync"
)

var wg sync.WaitGroup
var spiderName string = "Douban"

func Run() {
	urls := make([]string, 50)
	url := "https://www.bilibili.com"

	for i := 0; i < 50; i++ {
		urls[i] = url
	}

	downloader := NewDownloader(spiderName, &wg)
	spider := NewSpider(urls, spiderName, &wg)

	wg.Add(2)

	go spider.SpiderRun(urls, spider.StartParse)
	go downloader.Download()

	wg.Wait()

	defer close(ChRequest)
	defer close(ChResponse)

	log.Println("******DoubanSpider end******")
}

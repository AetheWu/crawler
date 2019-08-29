package crawler

import (
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type Downloader interface {
	Download() error
}

type DoubanDownloader struct {
	wg *sync.WaitGroup
}

func NewDownloader(spiderName string, wg *sync.WaitGroup) Downloader {
	var downloader Downloader
	switch spiderName {
	case "Douban":
		downloader = &DoubanDownloader{wg}
	}
	return downloader
}

func (d *DoubanDownloader) Download() error {
	log.Println("---func: Downloader start---")
	defer d.wg.Done()

	var wgi sync.WaitGroup
	defer wgi.Wait()

	i := 1

	for {
		select {
		case req := <-ChRequest:
			if req.callback != nil {
				for _, url := range req.urls {
					log.Println(i, "get:", url)
					wgi.Add(1)
					go d.get(url, req.callback, &wgi)
					i++
				}
			}

		case <-time.After(2 * time.Second):
			return nil
		}

	}

}

func (d *DoubanDownloader) get(url string, callback func(*Response, *sync.WaitGroup), wgi *sync.WaitGroup) {
	wgi.Done()
	log.Println("---func: get start---")

	resp, err := http.Get(url)

	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return
	}

	node, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		return
	}

	ChResponse <- &Response{node, callback}
	return
}

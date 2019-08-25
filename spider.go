package crawler

import (
	"log"
	"sync"
	"time"
)

type SpiderInterface interface {
	StartParse(*Response, *sync.WaitGroup)
	SpiderRun([]string, func(*Response, *sync.WaitGroup))
}

type Spiderx struct {
	urls []string
	name string
	wg   *sync.WaitGroup
}

func NewSpider(urls []string, name string, wg *sync.WaitGroup) SpiderInterface {
	var spider SpiderInterface

	src := Spiderx{urls, name, wg}

	switch name {
	case "Douban":
		log.Println("******DoubanSpider start******")
		spider = &DSpider{src}
	}

	return spider
}

func (s *Spiderx) InitSpider(urls []string, callback func(*Response, *sync.WaitGroup)) {
	if urls == nil {
		log.Println("start urls is null")
		return
	}

	log.Println("StartCrawl: ", urls)
	ChRequest <- &Request{urls, callback}
}

// func (s *Spiderx) StartParse(resp *Response, wg *sync.WaitGroup) {
// 	log.Println("---func StartParse have no implement---")
// 	return
// }

func (s *Spiderx) SpiderRun(urls []string, StartParse func(*Response, *sync.WaitGroup)) {
	log.Println("---func: SpiderRun start---")
	var wgi sync.WaitGroup
	defer log.Println("---func: SpiderRun end---")
	defer s.wg.Done()
	defer wgi.Wait()

	s.InitSpider(urls, StartParse)

loop:
	for {
		select {
		case v := <-ChResponse:
			log.Println("---callback start---")
			if v.callback != nil {
				wgi.Add(1)
				go v.callback(v, &wgi)
			}

		case <-time.After(2 * time.Second):
			break loop
		}
	}

}

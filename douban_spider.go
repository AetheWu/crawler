package crawler

import (
	"fmt"
	"sync"

	"golang.org/x/net/html"
)

type DSpider struct {
	Spiderx
}

func (s *DSpider) StartParse(v *Response, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("parse:", v.node.Data)
	var urls []string

	var parse func(*html.Node)

	parse = func(v *html.Node) {

		if v == nil {
			// log.Println("html is nil")
			return
		}

		if v.Type == html.ElementNode && v.Data == "a" {

			// v = v.FirstChild

			for _, attr := range v.Attr {
				if attr.Key == "href" {
					urls = append(urls, attr.Val)
				}
			}
		}
		parse(v.NextSibling)
		parse(v.FirstChild)

	}

	parse(v.node)
	fmt.Println("----result:", urls)
	fmt.Println("----")

	ChRequest <- &Request{urls, nil}

	// url := "https://movie.douban.com/"
	// ChRequest <- &Request{url, s.StartParse}
}

func (s *DSpider) parseMovie(v *Response, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("parse:", v.node.Data)

	var parse func(*html.Node)
	var urls []string

	parse = func(v *html.Node) {

		if v == nil {
			return
		}

		if v.Type == html.ElementNode && v.Data == "a" && v.Parent.Data == "div" {
			var href string

			for _, attr := range v.Attr {
				if attr.Key == "href" {
					urls = append(urls, attr.Val)
				}
			}

			if href != "" {
				fmt.Println(href)
			}
		}
		parse(v.NextSibling)
		parse(v.FirstChild)
	}

	parse(v.node)

	fmt.Println("----result:", urls)
	fmt.Println("----")

}

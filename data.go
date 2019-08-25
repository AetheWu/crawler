package crawler

import (
	"sync"

	"golang.org/x/net/html"
)

var ChRequest chan *Request = make(chan *Request, MAXLINKS)
var ChResponse chan *Response = make(chan *Response, MAXLINKS)

type Movie struct {
	title     string
	rate      string
	region    string
	directors []string
	actors    []string
}

type Request struct {
	urls     []string
	callback func(*Response, *sync.WaitGroup)
}

type Response struct {
	node     *html.Node
	callback func(*Response, *sync.WaitGroup)
}

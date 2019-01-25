package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]string
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Check(key string) bool {

	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	res := false
	//fmt.Println("->", key)

	if _, ok := c.v[key]; !ok {
		res = true
		//fmt.Println(val)
		c.v[key] = "val"
	}
	c.mux.Unlock()
	return res
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *SafeCounter) Crawl(ch chan string, url string, depth int, fetcher Fetcher) {
	//(url string, depth int, fetcher Fetcher) {

	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {

		if c.Check(url + body) {
			//fmt.Println(err)
			ch <- ("not found: " + url + body)
		}

		return
	}

	if c.Check(url + body) {
		//fmt.Printf("found: %s %q\n", url, body)
		ch <- ("found: " + url + body)
	}

	for _, u := range urls {
		go c.Crawl(ch, u, depth-1, fetcher)
	}

	return
}

func (c *SafeCounter) _Crawl(ch chan string, url string, depth int, fetcher Fetcher) {
	go c.Crawl(ch, url, depth, fetcher)
	//close(ch)
}

func (c *SafeCounter) PreCrawl(url string, depth int, fetcher Fetcher) {
	ch := make(chan string)
	go c._Crawl(ch, url, depth, fetcher)
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	c := SafeCounter{v: make(map[string]string)}
	//c.Crawl("https://golang.org/", 4, fetcher)
	c.PreCrawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

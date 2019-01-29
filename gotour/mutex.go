package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeMap struct {
	v   map[string]int
	mux sync.Mutex
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func (s *SafeMap) Checker(key string) bool {
	s.mux.Lock()
	res := false
	_, ok := s.v[key]
	if !ok {
		res = true
		s.v[key] = 1
	}
	if ok {
		s.v[key]++
	}
	s.mux.Unlock()
	return res

}

func Printer(url string, body string) {
	fmt.Printf("found: %s %q\n", url, body)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (s *SafeMap) Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {

		e := err.Error()
		if s.Checker(e) {
			fmt.Println(e)
		}
		return
	}

	//	fmt.Printf("found: %s %q\n", url, body)
	if s.Checker(url + body) {
		Printer(url, body)
	}

	for _, u := range urls {
		s.Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	s := SafeMap{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go s.Crawl("https://golang.org/", 4, fetcher)
	}

	time.Sleep(3 * time.Second)
	fmt.Println()
	for key, val := range s.v {
		fmt.Println(key, val)
	}
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

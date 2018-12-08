/*
In this exercise you'll use Go's concurrency features to parallelize
a web crawler.

Modify the Crawl function to fetch URLs in parallel without fetching
the same URL twice.

Hint: you can keep a cache of the URLs that have been fetched on a map,
but maps alone are not safe for concurrent use!
*/
package main

import (
	"time"
	"sync"
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type ParallelCrawler struct {
	fetched	map[string]int
	mux		sync.Mutex
}

func (c *ParallelCrawler) Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	c.mux.Lock()
	_, exists := c.fetched[url]
	if exists {
		c.mux.Unlock()
		return
	}
	c.fetched[url] = 0
	c.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go c.Crawl(u, depth-1, fetcher)
	}
	return
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	c := ParallelCrawler{
		map[string]int{},
		sync.Mutex{},
	}
	go c.Crawl(url, depth, fetcher)
	time.Sleep(time.Second)
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
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
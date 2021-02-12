package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type UrlCache struct {
	mu sync.Mutex
	urls  map[string] bool
}

// Adds new url to cache.
func (c *UrlCache) Add(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.urls.
	c.urls[key] = true
	c.mu.Unlock()
}

// Contains url query.
func (c *UrlCache) Contains(key string) bool {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.urls.
	defer c.mu.Unlock()
	return c.urls[key]
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c *UrlCache) {
	if depth <= 0 {
		return
	}
	
	if c.Contains(url) == true {
		return
	}
	
	body, urls, err := fetcher.Fetch(url)
	
	c.Add(url)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Printf("found: %s %q\n", url, body)
	
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, c)
	}
	return
}

func main() {
	c := UrlCache{urls: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, &c)
	time.Sleep(5*time.Second)
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

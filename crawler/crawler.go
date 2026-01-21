package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func crawl(url string, depth int, fetcher Fetcher, m map[string]bool, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Crawling:", url)
	fmt.Println("Found body:", body)
	//fmt.Println("urls:")
	for _, u := range urls {
		mutex.Lock()
		// fmt.Println(u)
		if !m[u] {
			m[u] = true
			mutex.Unlock()
			wg.Add(1)
			go crawl(u, depth-1, fetcher, m, mutex, wg)
		} else {
			mutex.Unlock()
		}
	}
}

func main() {
	m := map[string]bool{"https://golang.org/": true}
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go crawl("https://golang.org/", 4, fetcher, m, &mutex, &wg)
	wg.Wait()
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

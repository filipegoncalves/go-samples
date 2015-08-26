
package main

import (
	"fmt"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, ch chan string, visited_ch chan map[string]bool) {

	defer close(ch)

	if depth <= 0 {
		return
	}

	visited := <- visited_ch
	_, found := visited[url]
	visited[url] = true
	visited_ch <- visited

	if found {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}

	ch <- fmt.Sprintf("found: %s %q\n", url, body)

	chans := make([]chan string, len(urls))
	for i, u := range urls {
		chans[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, chans[i], visited_ch)
	}

	for i := range chans {
		for s := range chans[i] {
			ch <- s
		}
	}

	return
}

func main() {
	ch := make(chan string)

	visited_ch := make(chan map[string]bool, 1)
	visited_ch <- make(map[string]bool)

	go Crawl("http://golang.org/", 4, fetcher, ch, visited_ch)

	for s := range ch {
		fmt.Print(s)
	}
}

// fakeFetcher is a Fetcher that returns canned results.
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

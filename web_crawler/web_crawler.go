/* Taken from the Go Tour (Exercise: Web Crawler)
 *
 * In this exercise you'll use Go's concurrency features to parallelize a web crawler.
 *
 * Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.
 *
 */

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

	/* This implements an atomic "test and set" for the visited map. It atomically fetches the
         * visited status for the URL and sets it.
         *
         * This is cleverly achieved by using a buffered channel with unitary capacity where worker
         * threads consume the map when they want to read and mutate it, and write it back to the
         * channel once they're done.
         *
         * Note that the channel must be buffered with a capacity of 1, otherwise we would deadlock
         * because unbuffered channels block readers and writers until the other end is ready.
         *
         * This is Go's philosophy of concurrency:
         * Don't communicate by sharing memory, share memory by communicating
         *
         * How brilliant is that?
         */
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
		chans[i] = make(chan string, 128)
		go Crawl(u, depth-1, fetcher, chans[i], visited_ch)
	}

	/* This is how we implement synchronization and wait for other threads to finish.
         *
         * Each Crawl() thread is assigned its own channel to write results to. Each thread closes
         * its channel once it's done, that is, after writing its own results into the channel and
         * the results of the goroutines it spawned. Thus, results flow from a set of channels up
         * the "channel tree" until they reach the main, primary channel.
         *
         * This clever mechanism allows goroutines to wait for other spawned routines to terminate
         * before returning and closing their own channel.
         *
         * The channels are buffered (with a capacity of 128) because otherwise there is not much
         * parallelism, since each thread could only make progress after the parent thread fetched
         * the last result written.
         *
         * Synchronization is implicitly achieved with the channels, because each thread defers
         * closing the channel, which is wonderful.
         */

	for i := range chans {
		for s := range chans[i] {
			ch <- s
		}
	}

	return
}

func main() {
	ch := make(chan string, 128)

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

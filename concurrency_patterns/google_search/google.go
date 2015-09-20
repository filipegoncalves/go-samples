
package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Result string
type Search func(query string) Result

var (
	Web = []Search{
		fakeSearch("web1.google.com"),
		fakeSearch("web2.google.com"),
		fakeSearch("web3.google.com")}

	Image = []Search{
		fakeSearch("image1.google.com"),
		fakeSearch("image2.google.com"),
		fakeSearch("image3.google.com")}

	Video = []Search{
		fakeSearch("video1.google.com"),
		fakeSearch("video2.google.com"),
		fakeSearch("video3.google.com")}
)

func fakeSearch(kind string) func(query string) Result {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)

	for _, r := range replicas {
		go func() { c <- r(query) }()
	}

	return <- c
}

func Google(query string) (results []Result) {
	c := make(chan Result)

	go func() { c <- First(query, Web...) }()
	go func() { c <- First(query, Image...) }()
	go func() { c <- First(query, Video...) }()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case r := <- c:
			results = append(results, r)
		case <- timeout:
			fmt.Println("timed out")
			return
		}
	}

	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

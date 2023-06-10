package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type PageFetcher func(url string) ([]byte, error)

type ResultPrinter func(url string, result []byte)

// ProcessURLs processes a list of URLs with a specified number of goroutines
func ProcessURLs(urls []string, maxRoutines int, fetcher PageFetcher, printer ResultPrinter) {
	routines := make(chan int, maxRoutines)
	wg := sync.WaitGroup{}

	for _, url := range urls {
		routines <- 1
		wg.Add(1)
		go func(url string) {
			defer func() {
				<-routines
				wg.Done()
			}()

			page, err := fetcher(url)
			if err != nil {
				log.Println(err)
				return
			}

			printer(url, page)
		}(url)
	}

	wg.Wait()
}

// FetchPage fetches a page and returns its contents as a byte slice
func FetchPage(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching web page, url=%s: %w", url, err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body, url=%s: %w", url, err)
	}

	return body, nil
}

// PrintResult prints a URL and the MD5 hash of the page
func PrintResult(url string, result []byte) {
	fmt.Printf("%s %s\n", url, GetMD5Hash(string(result)))
}

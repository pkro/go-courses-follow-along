package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func sitesSerial(urls []string) {
	for _, url := range urls {
		returnType(url)
	}
}

func sitesConcurrent(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		// add 1 to the number of goroutines to wait for
		wg.Add(1)
		// go runs the function (here an anonymous IIFE, but could be any function)
		// in a separate goroutine
		go func(url string) {
			returnType(url)
			wg.Done() // removes 1 from waitgroup
		}(url)
		wg.Wait() // waits for all the goroutines to finish
	}
}
func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}

	start := time.Now()
	sitesSerial(urls)
	fmt.Println(time.Since(start))

	start = time.Now()
	sitesConcurrent(urls)
	fmt.Println(time.Since(start))
}

package main

import (
	"context"
	"fmt"
	"time"
)

type Bid struct {
	AdURL string
	Price float64
}

func bestBid(url string) Bid {
	time.Sleep(20 * time.Millisecond)
	return Bid{"http://example.com", 1.99}
}

var defaultBid = Bid{"http://example.com/default", 0.5}

func findBid(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1) // buffered to avoid goroutine leak
	go func() {
		ch <- bestBid(url)
	}()

	select {
	case bid := <-ch:
		return bid

	case <-ctx.Done():
		return defaultBid
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	url := "https://http.cat/418"
	bid := findBid(ctx, url)
	fmt.Println(bid) // {http://example.com 1.99}

	// lower timeout
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	url = "https://http.cat/404"
	bid = findBid(ctx, url)
	fmt.Println(bid) // {http://example.com/default 0.5}
}

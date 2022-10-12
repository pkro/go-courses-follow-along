package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch1 <- 42
	}()

	// once one of the channels becomes available (true for receving AND sending),
	// it receives from the channel of the selected case
	select {
	case val := <-ch1:
		fmt.Printf("got %d from ch1\n", val)
	case val := <-ch2:
		fmt.Printf("got %d from ch2\n", val)
	}

	// usage for timouts

	out := make(chan float64)
	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Printf("got %f\n", val)
	// time.After runs another goroutine behind the scenes
	// and then sends a signal over a channel
	case val := <-time.After(20 * time.Millisecond):
		fmt.Printf("timeout signal %v", val) // timeout signal 2022-10-12 07:28:14.529286716 +0200 CEST m=+0.020190679
	}
}

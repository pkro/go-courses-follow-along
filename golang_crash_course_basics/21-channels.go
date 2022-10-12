package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// goroutine runs concurrently to the main one
	go func() {
		ch <- 353
	}()

	val := <-ch // receive
	fmt.Println(val)

	// send multiple
	const count = 3
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		// shuts down the channel after the last value is received
		// must be only done by the sender
		close(ch)
	}()

	for i := 0; i < count; i++ {
		val := <-ch
		fmt.Printf("Received %d\n", val)
	}

	/* output:
	sending 0
	Received 0
	sending 1
	Received 1
	sending 2
	Received 2
	*/

	// buffered channels
	ch2 := make(chan int, 1)
	ch2 <- 19
	// does NOT block at this point as the channel is buffered now
	val2 := <-ch2

	fmt.Println(val2)

}

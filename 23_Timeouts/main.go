package main

import (
	"fmt"
	"time"
)

/*
Timeouts are important for programs that connect to external resources tor that otherwise need to bound execution time. Implementing timeouts in Go is easy and elegant thanks to channels and select.
*/

func main() {

	/*
		For our example ,suppose we're executing and external call that returns it's result on a channel c1 after 2s.
	*/
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	/*
		Here's the select implementing a timeout. res := <-c1 awaits the result and <- Time.After wawauts a value to be sent after the timeout of 1s. Since select proceeds with the first recive that's ready, we'll take the timout case if the operation takes more than the allwoed 1s.

		If we allow a longer timout of 3s, then the receive from c2 will succeed and wel'' print the result.
	*/
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

	/*

		Running this program shows the first operation timing out and the second succeding.

		Using the select timeout pattern requries communicating results over channels. This is a good idea in gernral because other important Go features based on channels and select. We'll look at two examples of this next: timers and tickers.
	*/
}

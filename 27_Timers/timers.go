/*
We often want to execute Go code at some point in the future, or repeatedly at some interval. Go's built-in timer and ticker features make both of these tasks easy. We'll look first at timers and then at tickers.
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		Timer represent a single event in the future. You tell the timer how long you
		want to wait, and it provides a channel that wil be notified at that time.
		This timer will wait 2 seconds.
	*/
	timer1 := time.NewTimer(time.Second * 2)

	/*
		The <-timer1.C blocks on the timer's channel C until it sends a value indicatiing that the timer expired.
	*/
	<-timer1.C
	fmt.Println("Timer 1 expired")

	/*
		If you just wanted to wait, you could have used time.Sleep. One reaosn a timer may be useful is that you can cancel the timer before it expries. Here's an example of that.
	*/
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	/*
		The first timer will expire ~2s after we start the program, but the second should be stopped before it has a chance to expire.
	*/
}

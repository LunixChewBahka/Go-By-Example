/*
In this example we'll look at how to implement a worker pool using goroutines
and channels
*/
package main

import (
	"fmt"
	"time"
)

/* Here's the worker, of which we'll run several concurrent instances. These wokrers will receive work on the jobs channel and send the corresponding results on results. We'll sleep a per job to simulate an expensive task.
 */
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}
}

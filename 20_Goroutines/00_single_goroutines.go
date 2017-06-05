package main

import (
	"fmt"
	"time"
)

// benchmark function using testing package
func worker(finished chan bool) {
	fmt.Println("Worker: Started")
	time.Sleep(time.Second * 5)
	fmt.Println("Worker: Finished")
	finished <- true
}

func main() {

	start := time.Now()
	finished := make(chan bool)

	fmt.Println("Main: Starting worker")
	elapsed_00 := time.Since(start)
	fmt.Println(elapsed_00)

	go worker(finished)

	elapsed_01 := time.Since(start)
	fmt.Println(elapsed_01)

	fmt.Println("Main: Waiting for worker to finish")
	<-finished
	elapsed_02 := time.Since(start)
	fmt.Println(elapsed_02)
	fmt.Println("Main: Completed")

	elapsed_03 := time.Since(start)
	fmt.Println(elapsed_03)
}

package main

/*

https://gobyexample.com/worker-pools

In this example we'll look at how to implement a worker
pool using goroutines and channels.

Here's the worker, of wich we'll run several concurrent instances.
These workers will receive work on the jobs channel and send the
corresponding results on results. We'll sleep a second per job to
simulate an expensive task.

In ordeer to use our pool of workers we need to send them work and
collect their results. We make 2 channels for this.

This starts up 3 workers, initially blocked because there are no jobs yet.

Here we send 5 jobs and then close that channel to indicate that's all the work we have.

Finally we collect all the results of the work.

Our running program shows that 5 jobs being executed by various workers.
The program only takes about 2 seconds despite doing about 5 seconds of total work
because there are 3 workers operation concurrently.

*/

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
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

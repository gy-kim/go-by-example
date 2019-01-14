package main

/*

https://gobyexample.com/atomic-counters

The primary mechanism for managing state in Go is communication over channels.
We saw this for example with worker pools. There are a few other options for
managing state though. Here we'll look at using the sync/atomic package
for atomic couters accessed by multiple goroutines.

*/
import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	/*
		We'll use an unsigned integer to represent our (always positive) couter.
	*/
	var ops uint64

	/*
		To simulate concurrent updates, we'll start 50 goroutines
		that each increment the couter about once a millisecond.
	*/
	for i := 0; i < 50; i++ {
		go func() {
			for {
				/*
					To atomically increment the couter we use AddUint64,
					giving it the memory address of our ops couter with the &syntax.
				*/
				atomic.AddUint64(&ops, 1)

				// Wait a bit between increments.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Wait a second to allow some ops to accumulate.
	time.Sleep(time.Second)

	/*

		In order to safety use the counter while it's still being updated by other gorines,
		we extract a copy of the current value into opsFinal via LoadUint64.
		As above we need to give this function the memory address &ops from which to fetch the value.

	*/
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
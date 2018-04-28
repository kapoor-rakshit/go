// Concurrency is dealing with multiple things at once, 
// and it is not actually doing multiple things simultaneously (parallelism).

// A Go Routine is a lightweight thread, which we can easily spawn by simply prefacing the function call with keyword 'go'
// goroutines aren't going to necessarily run or complete before program ends.

package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup                      // define the WaitGroup variable

func say(s string) {
	
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
}

func wgsay(s string) {

		defer wg.Done()
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
}

func main() {
	go say("Hey")     // a goroutine
	say("there")      // normal func call

	fmt.Println()

	say("Hey")           // normal func call
	go say("there")      // a goroutine
// normal func is going to run in full first, hogging resources for itself first. 
// Unfortunately, after this, the only other line of code is for the goroutine to run, 
// which doesn't really have any priority to actually keep the program overall running,

    fmt.Println()

	go say("Hey")         // a goroutine
	go say("there")      // a goroutine
	time.Sleep(1000*time.Millisecond)        // without this statement, no output is visible
// this only works if we know the exact time that our goroutines will take, or something close. 
// If we're too high in our estimation, we waste time, if we're too low, our processes don't run! 

	fmt.Println()

// there is a problem with above methods where goroutines will not necessarily finish by the time program does.
// We need a way to wait for them as necessary.

// Besides waiting for goroutines to finish, where they aren't dependent on eachother, 
// there are also going to be times when we want to run a group of goroutines, wait for them to all finish, then run another group. 
// Again, we need some way, other than an arbitrary sleep, to make this happen. To do it, we'll use the sync package from the standard library.	

// Add() : identify how many goroutines need to be waited.
// Done(): When a goroutine exits, it must call Done.
// Wait(): 'main' goroutine blocks on Wait, Once the counter becomes 0, the Wait will return, and main goroutine can continue to run.

	wg.Add(1)                // add goroutines to this group
	go wgsay("Hey")
	wg.Add(1)
	go wgsay("there")
	wg.Wait()                // wait for this group of goroutines with wg.Wait()

	fmt.Println("This is main goroutine")
} 
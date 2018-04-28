// The idea of 'panic' is to halt the program and start to panic.
// This will stop the running function and run all of the deferred functions from the panicking function.

// The 'recover' function lets you recover from a panicking goroutine.
// To recover a panicking goroutine, you would need to use recover from within one of the deferred functions of that goroutine.

package main

import (
	"fmt"
)

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup:", r)
	}
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		if i == 2 {
			panic("2 encountered")
		}
		fmt.Println(s)
	}	
}

func main() {
	say("Hey")
}
// Go Channels are used to connect concurrent goroutines, to send and receive values between them, using channel operator '<-'

// Channels by default are blocking on sending and receiving values, so they will be waited on. 

// MORE : https://pythonprogramming.net/go/channel-buffer-range-go-language-programming-tutorial/
// Go Channels buffering, iteration, and synchronization

package main

import 
(
	"fmt"
)

func foo(c chan int, someValue int) {
    c <- someValue * 5                      // send value over channel
}

func main() {

    fooVal := make(chan int)              // create a channel, using the 'make(chan datatype)' built-in function

    go foo(fooVal, 5)                     // goroutines that are sending some values
    go foo(fooVal, 3)

    v1 := <- fooVal                       // receive value from channel
    v2 := <- fooVal

    fmt.Println(v1, v2)
}
/* defer statement is used to put off (defer) the execution of something until the surrounding function is done.
The defer statement will be evalutated immediately, but won't run until the end, so we usually put the defer statement early in the function.
The main idea here is to do this to ensure things like files are closed and connections are terminated.  */

/* The defer statement works on a "first in last out" structure, so the last defer will also be the first one that runs */

package main

import "fmt"

func foo() {
  defer fmt.Println("Done!")                             // 3
  defer fmt.Println("Are we done yet?")                  // 2
  fmt.Println("Doing some stuff, who knows what?")       // 1
}

func main() {
  foo()
}
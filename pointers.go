package main

import (
	"fmt"
)

func main() {
	a := 100
	x := &a

	fmt.Println(x,&a)
	fmt.Println(*x,a)

	*x = *x * *x                    // note use of = and not :=
	fmt.Println(*x,a)

	fmt.Println(x,&a)
}
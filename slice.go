
// Go Slice is an abstraction over Go Array. or Similar to list[] in Python.

package main

import (
	"fmt"
)

func main() {
	var a[] int                             // slice is declared as an array without size
	var tp int

	for i:=0;i<5;i++ {
		fmt.Scanf("%d",&tp)
		a = append(a,tp)                   // append values to slice
	}

	fmt.Println(a,len(a),cap(a))           // len() returns the elements presents in the slice
	                                   // cap() returns the capacity of the slice (i.e. how many elements it can be accommodate)


	aa := make([]int,0,7)                // empty slice      // make(type, len, cap)
	                                    // len=0 gives an empty slice, else slice with 0's created [0,0,0,...]
	fmt.Println(aa,len(aa),cap(aa))
	aa = append(aa,2,7,8)
	aa[2] = 9025
	fmt.Println(aa,len(aa),cap(aa))

	newaa := aa[:2]                           // new slice created with subslicing operations
	fmt.Println(newaa,len(newaa),cap(newaa))

	newa := make([]int,len(a)*2,cap(a)*2)       // params to make()
	copy(newa,a)                              // copy a slice with same params of type
	fmt.Println(newa,len(newa),cap(newa))
}
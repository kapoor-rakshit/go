package main

import(
"fmt"
)

func testfunc(val float64){                                // testfunc(name type, name type, ...) return type { func body }
	fmt.Printf("Testing func def and call.\n")
	fmt.Println("Param passed",val)
}

func testret(val1,val2 int, vala,valb float64) (float64,int) {    // multiple returns
	return vala+valb , val1+val2
}

func main() {
	testfunc(5.258)

	fsum , isum := testret(2,4,6.39,5.0)

	fmt.Println(fsum,isum) 
}
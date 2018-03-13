
// Use errors.New() to construct a basic error message

package main

import (
	"fmt"
	"math"
	"errors"                                   // package for error handling
)

func getsqrt(n float64) (float64,error) {     // return type "error"

	if n<0 {
		return 0, errors.New("negative number passed")  // generate error message
	}else {
		return math.Sqrt(n),  nil                      // nil === NULL
	}
}

func main() {

	res, err := getsqrt(-66)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(res)
	}

	res, err = getsqrt(36)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(res)
	}
}
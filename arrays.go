package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func testarrpass(passedarr [] int){           // array as an arg
	l := len(passedarr)
	fmt.Println(l)
	fmt.Println(passedarr[2])
}

func main() {

	var i,j int

	// var arrval[5] int                       // 1-D array declaration, only const size is valid

	var arrval = [5] int {4,5,6,7,9}          // values defined for fixed size
	fmt.Println(len(arrval))

	var arrdyn = [] int {33,6,78}             // values defined, size is assigned dynamically on no. of values
	fmt.Println(len(arrdyn))

	reader := bufio.NewReader(os.Stdin)      
	vals,_ := reader.ReadString('\n')       // space separated values as input

	sep_newline_char := strings.Split(vals,"\n")                // values splitted at delim "\n"
	vals_without_newlinechar := sep_newline_char[0]
	arr := strings.Split(vals_without_newlinechar," ")          // values splitted at delim " "

	val1,_ := strconv.ParseInt(arr[0],10,64)
	val2,_ := strconv.ParseInt(arr[1],10,64)                   // Parse to Int64    (val, base10, Int64)
	fmt.Println(val1,val2)


                                               /* 2D array */
	var matval = [2][2] int {{1,2} , {55,6}}
	fmt.Println(matval[1][0])

	var matdyn = [][] int {{45,6} , {15,26,8}}   // diff column sizes
	fmt.Println(matdyn[0][1])

	var mat[2][3] string                        // 2-D array declaration
	for i=0;i<2;i++ {
		row,_ := reader.ReadString('\n')

		temp := strings.Split(row,"\n")
		rowf := temp[0]
		rowarr := strings.Split(rowf," ")
		for j=0;j<3;j++ {
			mat[i][j] = rowarr[j]              // definition
		}
	}

	fmt.Println(mat[0][2])

	testarrpass(arrdyn)                       // pass array in a func
}
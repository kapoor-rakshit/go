
/* In the Go language, there is only the "for" loop, there is no "while" statement */

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	n,_ := reader.ReadString('\n')
	tp := strings.Split(n,"\r")                     // remove last "\r" from args
	l,_ := strconv.Atoi(tp[0])

	str,_ := reader.ReadString('\n')
	tp = strings.Split(str,"\r")
	arr := strings.Split(tp[0], " ")
	
	for i:=0 ; i<l ; i++ {                         // for loop
		fmt.Printf("%s ",arr[i])
	}

	i := 0                                         // i is dynamically typed, as prev version scope was within previous loop
	                                              //  this i is available for main func and modified acc.

	for i<2 {                                      // a mimic of "while" loop
		fmt.Print("hey"," ")
		i+=1
	}

	for {                                          // Infinite loop, a mimic of "while(True)"
		if i>=5 {
			break
		}
		fmt.Print(i," ")
		i+=1
	}

	var temp[] int                               // empty slice
	var v int
	for i:=0;i<4;i++ {
		fmt.Scanf("%d",&v)
		temp = append(temp,v)
	}
	for i:= range temp {                      // iterate slice using "range" keyword, which gives index
		fmt.Print(temp[i]," ")
	}
}
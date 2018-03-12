// Reference    : https://pythonprogramming.net/go/introduction-go-language-programming-tutorial/
// installation : https://golang.org/doc/install     
//                installer sets env vars itself 
// to run prog  : go run filename.go

// With the Go programming language, we can't define variables that we won't use

package main        // main package

                   /* import libs */
import ("fmt";"math")                                       // fmt lib - formatted I/O
import (
	"time"
	"math/rand"                                             // package within package
	"bufio"
	"os"
	"strings"
)

var a,b int = 32,56                                         // dynamic not allowed outside func, so statically typed
var c float64

func main() {
	fmt.Println("First Go prog")                            // Println()
	fmt.Printf("%f\n%f\n",math.Sqrt(16), math.Pow(5,3))     // Printf()     //math lib

	fmt.Println("Time is",time.Now())                       // time()

	fmt.Println(rand.Intn(12));                             // rand() returns [0,n)

	x,y := 7.89, 4.0                                       // dynamic typing inside functions allowed, using :=
	fmt.Println(int(x/y))                                  // typecasting

	var d bool                                             // static typing
	d = true

	c = x                                                 // since c already defined so statically assigned

	str := "Hey golang"
	
	fmt.Println(a,b,c,d,str)                             // data types - int, float64, bool, string

	var sv1,sv2 int
	var st string
	fmt.Scanf("%d%d%s",&sv1,&sv2,&st)                   // read from stdin, SPACE separated values

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')                           // to prevent '\n' be read
	naam, _ := reader.ReadString('\n')                // desired string read until delim '\n' occur
	                                                  // else return error in _ 

	arr := strings.Split(naam," ")                    // split values with delimiter as " " and store in array

	fmt.Println(sv1,sv2,st)
	fmt.Print(naam)

	fmt.Println(arr[4])
}
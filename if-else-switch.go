
// Operators used are && , || , !

// No "break" is needed in the case statement of switch construct

package main

import (
	"fmt"
)

func main() {
	a := 50
                                          /* if-else construct */
	if a>15 && a>45 {
		fmt.Printf("Greater\n")
	}else if a<15 {                       // "else if" must be inline with terminating } of "if"
		fmt.Printf("Lesser\n")
	}else {                               // "else" inline with terminating } of "else if"
		fmt.Printf("Equal\n")
	}

                                         /* Switch construct */
	marks := 50
                                          // Example 1
	switch marks {                        // boolean-expression or integral type to "switch" and "case"
	case 25,50,75:
		fmt.Printf("yes matched 25m\n")
	case 12,24:                           // multiple params to "case" allowed
		fmt.Printf("yes matched 12m\n")
	default:
		fmt.Printf("No match found\n")
	}
                                           // Example 2
	switch {                               // If the expression is not passed to "switch" then the default value is true
	case marks == 12, marks == 24:         // boolean-expression or integral type to "case" and "switch"
		fmt.Printf("12m matched\n")
	case marks == 25:
		fmt.Printf("25m matched\n")
	default:
		fmt.Printf("No match\n")
	}
                                        // Example 3 ... type switch
	var x interface{}                   // expression used in a "switch" statement must have an variable of interface{} type.
	x = 32.56

	switch x.(type) {
	case nil:
		fmt.Printf("It is of %T type",x)
	case int:
		fmt.Printf("It is of %T type",x)
	case float64:
		fmt.Printf("It is of %T type",x)
	case func(int) float64:
		fmt.Printf("It is of %T type",x)
	case bool, string:
		fmt.Printf("It is of %T type",x)
	default:
		fmt.Printf("Unknown type\n")
	}
                                      /* Select construct */
	
	// Reference : https://www.tutorialspoint.com/go/go_select_statement.htm
	// The type for a "case" must be the a communication channel operation.
}
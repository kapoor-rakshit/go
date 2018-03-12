/* Go has no Classes. 
   Instead, it has structs, and we can have methods by defining methods on the structs. 

   Methods that just access values are called value receivers and 
   Methods that can modify information are pointer receivers.*/

package main

import (
	"fmt"
)

const conv_factor float64 = 10

type stud struct {

	roll_no int
	name string
	cgpa float64
}
                              /* A value receiver 
                                 We use s (variable) and then stud (type) which is in association with the stud struct. 
                                 The method gets a copy of the object, so we cannot actually modify it here */
func (s stud) getperc() float64 {

	return s.cgpa * conv_factor
}
                               /* A pointer receiver 
                                  Note the *stud in func (s *stud), we're modifying the struct itself via pointer. */
func (s *stud) updatetcgpa(cgpa float64) {
	
	s.cgpa = cgpa
}

func main() {

	var naam string
	fmt.Scanf("%s",&naam)

	studinst1 := stud{9025, naam, 8.42}                           // variables or hardcoded values

	studinst2 := stud{name : "Tak" , roll_no : 84 , cgpa : 0}      // other way to instantiate

	fmt.Println("Stud1Name :",studinst1.name,"\n","Stud1Roll :",studinst1.roll_no, "\n", studinst1.cgpa)
	fmt.Println("Stud2Name :",studinst2.name,"\n","Stud2Roll :",studinst2.roll_no, "\n", studinst2.cgpa)

	fmt.Println("Stud1Perc :",studinst1.getperc(),"%")
	fmt.Println("Stud2Perc :",studinst2.getperc(),"%")                 // use of value receiver

	studinst1.updatetcgpa(8.45)
	fmt.Println("Stud1Modified :",studinst1.cgpa, studinst1.getperc(), "%")     // use of pointer receiver
}
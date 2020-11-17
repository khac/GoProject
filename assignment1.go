// total marks: declare 5 variables for 5 subjects, inititalize marks for each of them, and print the total marks

package main

import (
	"fmt"
)

func main() {
	var sub1 int = 80
	var sub2 int = 81
	var sub3 int = 82
	var sub4 int = 83
	var sub5 int = 84
	fmt.Printf("Total: %d \n", sub1+sub2+sub3+sub4+sub5)
}

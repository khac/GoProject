// Write a function named "isEven" to check if a given integer and return true if it is even, otherwise false. 


package main

import (
	"fmt"
)

func main() {
	m := 1012
	n := 98981
	
	fmt.Println(isEven(m))		
	fmt.Println(isEven(n))
}

func isEven(s int) (n1 bool) {
	if(s%2==0){
		return true
	} else{
		return false
	}
}

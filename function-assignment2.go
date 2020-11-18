// write a function that accepts two input parameters and returns two value. 
// example a swap function


package main

import (
	"fmt"
)

func main() {
	var a int = 98
	var b int = 1221
	fmt.Println(a,b)
	
	a, b = swap(a, b)
	fmt.Println(a,b)
}

func swap(s int, t int) (int, int) {
	return t, s
}

// write a function that accepts two input parameters and returns two value. 
// example a swap function. 
// and call this function in main, and store the first return value, ignoring the second return value

package main

import (
	"fmt"
)

func main() {
	var a int = 98
	var b int = 1221
	fmt.Println(a,b)
	
	a, _ = swap(a, b)
	fmt.Println(a,b)
}

func swap(s int, t int) (int, int) {
	return t, s
}

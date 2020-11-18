// Write a function named "swap" to swap 2 given integers and return the swapped numbers

package main

import (
	"fmt"
)

func main() {
	m := 1012
	n := 9898
	
	m,n = swap(m,n)
	fmt.Println(m,n)
}

func swap(s int, t int) (n1 int, n2 int) {
	n1 = t
	n2 = s
	return
}

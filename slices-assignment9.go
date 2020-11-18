// Given a slice of numbers, search if a given number is in the slice or not

package main

import (
	"fmt"
)

func main() {
	set := []int{1, 2, 3, 4, 5}

	fmt.Println(find(set, 11))
}

func find(s []int, f int) bool {
	// add your code here
	for i := 0; i < len(s); i++ {
		if s[i] == f {
			return true
		}
	}
	return false
}

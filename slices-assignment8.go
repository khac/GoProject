// Create a slice of integer values, and loop thru all the values,
// and prepare another slice with only the odd numbers
// return nil if no match found
// no need to sort the data

package main

import (
	"fmt"
)

func main() {
	set := []int{12, 992, 100, 1012, 102}

	fmt.Println(oddSet(set))
}

func oddSet(s []int) []int {
	// add your code here
	var index int = 0
	var slice []int = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i]%2 == 1 {
			slice[index] = s[i]
			index = index + 1
		}
	}
	if index == 0 {
		return nil
	}
	return slice[:index]
}

// Create a slice of integer values, and loop thru all the values,
// and prepare another slice with only the even numbers
// return nil if no match found
// no need to sort the data

package main

import (
	"fmt"
)

func main() {
	set := []int{10, 99, 100, 101, 102}

	fmt.Println(evenSet(set))
}

func evenSet(s []int) []int {
	// add your code here
	slice := make([]int, len(s))
	var index int = 0
	for i := 0; i < len(s); i++ {
		if s[i]%2 == 0 {
			slice[index] = s[i]
			index = index + 1
		}
	}
	return slice[:index]
}

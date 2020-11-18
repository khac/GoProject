// Given a slice of numbers, print the sum all the numbers

package main

import (
	"fmt"
)

func main() {
	set := []int{1, 2, 3, 4, 5}

	fmt.Println(sum(set))
}

func sum(s []int) int {
	// add your code here
	var sum int = 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}

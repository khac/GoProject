// Write a function named "sum" to input a slice of integers and return the sum of all its elements

package main

import (
	"fmt"
)

func main() {
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	fmt.Println(sum(m))
}

func sum(s []int) (cur int) {

	for _, v := range s {
		cur += v
	}
	return cur
}

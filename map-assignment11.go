// Given a slice of numbers, eliminate duplicates and print the output
// if empty or nil, return as is
// output need not be in any specific order

package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 1, 1, 2, 1, 1, 2, 1, 3}

	fmt.Println(dedup(s))
}

func dedup(s []int) []int {
	// add your code here
	mapofArr := map[int]bool{}
	index := 0
	var slice []int = make([]int, len(s))
	for _, i := range s {
		if _, ok := mapofArr[i]; ok {

		} else {
			mapofArr[i] = true
			slice[index] = i
			index = index + 1
		}
	}

	return slice[:index]

}

// given a slice of int, find the largest number

package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 44, 5, 6, 7, 8, 9}
	fmt.Println(largest(s))
}

func largest(s []int) int {
	// add code here
	var max int = s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max

}

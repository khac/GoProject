// write a function named "isSubset", given 2 slices of int, find if the second slice is subset of first one. return true if it is a subset, otherwise false

package main

import (
	"fmt"
)

func main() {
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	n := []int{1, 2, 3, 4}
	o := []int{1, 2, 3, 4, 123, 1233}

	fmt.Println(isSubset(m, n))
	fmt.Println(isSubset(m, o))

}

func isSubset(s []int, t []int) bool {
	m1 := make(map[int]bool)

	for _, v := range s {
		m1[v] = true
	}

	for _, v := range t {
		if _, ok := m1[v]; ok {
			continue
		} else {
			return false
		}
	}

	return true
}

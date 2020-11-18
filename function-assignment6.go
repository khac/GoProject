// Write a function name "splitEvenOdd" to input a slice of integers, and
// return two slices of integers. first one with all the even numbers from the input slice, and second one with the odd numbers

package main

import (
	"fmt"
)

func main() {
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	a, b := funcMap(m)
	fmt.Println(a)
	fmt.Println(b)
}

func funcMap(s []int) (s1 []int, s2 []int) {
	s1 = make([]int, len(s))
	s2 = make([]int, len(s))
	var indexOdd int
	var indexEven int
	for _, v := range s {
		if v%2 == 1 {
			s1[indexOdd] = v
			indexOdd += 1
		} else {
			s2[indexEven] = v
			indexEven += 1
		}
	}
	s1 = s1[:indexOdd]
	s2 = s2[:indexEven]
	return
}

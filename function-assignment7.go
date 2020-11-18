// Write a function named "concat" to input a slice of strings, and return a concatenated string of all its elements

package main

import (
	"fmt"
)

func main() {
	m := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	stringCur := concat(m)
	fmt.Println(stringCur)
}

func concat(s []string) (cur string) {

	for _, v := range s {
		cur += v
	}
	return cur
}

// Create a slice of int, initialize it, and print the last element of the slice.
// return 0 if empty

package main

import (
	"fmt"
)

func main() {
	sl := []int{}
	fmt.Println(lastElement(sl))
}

func lastElement(sl []int) int {
	// add your code here
	if len(sl) == 0 {
		return 0
	} else {
		return sl[len(sl)-1]
	}

}

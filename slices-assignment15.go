// Create a slice of int, initialize it, and print the first element of the slice
// return 0 if empty

package main

import (
	"fmt"
)

func main() {
	sl := []int{}
	fmt.Println(firstElement(sl))
}

func firstElement(sl []int) int {
	// add your code here
	if len(sl) == 0 {
		return 0
	}
	return sl[0]

}

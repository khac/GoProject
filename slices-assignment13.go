//declare and initialze a slice of int and find the length of the slice

package main

import (
	"fmt"
)

func main() {
	set := []int{1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 9} // TODO: initialze the slice

	fmt.Println(length(set))
}

func length(s []int) int {
	// TODO: return the length of the slice
	return len(s)

}

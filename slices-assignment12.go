//declare and initialze a slice of int and find the capacity of the slice

package main

import (
	"fmt"
)

func main() {
	set := []int{1,2,3,4,5,6,7,8,9,10} // TODO: initialze the slice
	fmt.Println(capacity(set))
}

func capacity(s []int) int {
	// TODO: return the capacity of the slice
	return cap(s);
}

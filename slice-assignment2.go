// Create a slice of bytes for a string value
// hint: string type is internally equivalent to byte slice.

package main

import (
	"fmt"
)

func main() {
	myslice := "goplayground"
 	secondSlice := myslice[1:4]

	fmt.Println(myslice)
	fmt.Println(secondSlice)
}

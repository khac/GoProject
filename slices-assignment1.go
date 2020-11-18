// Create 2 slices, initialize only one, and copy it to the another one, and print them

package main

import (
	"fmt"
)

func main() {
	var myslice []int
	var secondSlice []int
	myslice = []int{1, 2, 3, 4, 5, 6, 7}
	secondSlice = myslice
	fmt.Println(myslice)
	fmt.Println(secondSlice)
}

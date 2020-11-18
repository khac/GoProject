// Create a slice and loop over slice printing out both the index and the value

package main

import (
	"fmt"
)

func main() {
	myslice := "abcdefghijklmnopqrstuvwxyz"
 	secondSlice := myslice[1:8]

	for i,v := range secondSlice{
		fmt.Println(i,v)
	}
	
}

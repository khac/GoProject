// Create a slice, initialize it, and using for loop, print all the elements of the slice

package main

import (
	"fmt"
)

func main() {

 	slice := []byte{'a','b','c','d','e','f','g'}

	for i,v := range slice{
		fmt.Println(i,v)
	}
	
}

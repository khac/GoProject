// Create a slice of integer values, and loop thru all the values,
// and prepare another slice with only the values which are greater than 100.
// return nil if no match found
// no need to sort the data

package main

import (
	"fmt"
)

func main() {
	set := []int{1, 99, 100, 10, 10}

	fmt.Println(subset(set))
}

func subset(s []int) []int {
	// add your code here
	ret := make([]int, len(s))
	index := 0
	for i:=0; i<len(s); i++{
		if(s[i]>100){
			ret[index] = s[i]
			index = index+1
		}
	}
	if(index==0) {
		//fmt.Println("Here")
		return nil
	}
	return ret[0:index];
}

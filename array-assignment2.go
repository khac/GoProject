// sum of all elements in an array of int of size 4

package main

import (
	"fmt"
)

func main() {
	a := [4]int{1, 2, 3, 4}
	sum := arraySum(a)
	fmt.Println(sum)
}

func arraySum(a [4]int) (sum int) {
	// add code here:
	sum = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

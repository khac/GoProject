// given a map of int to int, find the sum of the values for which key is a multiple of 3
// return 0 when the map is nil or empty, or if no key is a multiple of 3

package main

import (
	"fmt"
)

func main() {
	marks := map[int]int{1: 90, 2: 32, 3: 80}

	fmt.Println(multiple3sum(marks))
}

func multiple3sum(m map[int]int) int {
	//TODO: write your code here
	sum := 0
	for k, v := range m {
		if k%3 == 0 {
			sum += v
		}
	}
	return sum
}

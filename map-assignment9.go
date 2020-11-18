// given a map of string to int, find the sum of all the (integer) values
// return 0 when the map is nil or empty

package main

import (
	"fmt"
)

func main() {
	//	marks := map[string]int{"math": 90, "science": 80}
	var marks1 map[string]int = nil

	fmt.Println(sum(marks1))
}

func sum(m map[string]int) int {
	//TODO: write your code here
	var sum int = 0
	for _, v := range m {
		sum += v
	}
	return sum

}

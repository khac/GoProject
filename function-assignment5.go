// create a function with a map (of int to int) as input and returns the same

package main

import (
	"fmt"
)

func main() {
	var m = make(map[int]int)
	fmt.Println(funcMap(m))
}

func funcMap(s map[int]int) map[int]int {
	s[1] = 111
	s[2] = 222
	s[3] = 333
	s[4] = 444
	s[5] = 555
	s[6] = 666

	return s
}

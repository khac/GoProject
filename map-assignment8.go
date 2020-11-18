//to check if the map contains the given key using comma ok idiom (val, exists)

package main

import "fmt"

func main() {

	myMap := map[int]string{
		0: "Zero",
		1: "One",
		2: "Two",
		3: "Three",
	}

	e := exists(myMap, 2)
	fmt.Println(e)
}

func exists(m map[int]string, k int) bool {
	// add your code here
	if _, ok := m[k]; ok {
		return true
	}
	return false
}

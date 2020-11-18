//to a map with key as string and values as int; assign the given key-value pair to that map

package main

import (
	"fmt"
)

func main() {
	m := map[string]int{}
	m2 := set(m, "hello", 100)
	fmt.Println(m2)
}

func set(m map[string]int, k string, v int) map[string]int {
	// add your code here
	m[k] = v
	return m
}

// Write a function named "add" to add two integers and return the sum

package main

import (
	"fmt"
)

func main() {
	m := 1018
	n := 98981

	fmt.Println(add(m, n))

}

func add(s int, t int) int {
	return s + t
}

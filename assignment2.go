// Assignment: Create two integer variables and assign the value of one variable to another

package main

import (
	"fmt"
)

func main() {
	var firstVar = 10
	var secondVar = firstVar
	fmt.Printf("first: %d, second: %d", firstVar, secondVar)
}

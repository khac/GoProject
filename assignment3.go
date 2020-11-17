// declare a constant to store the math pi value 3.414,
// and use it to calculate the area of a circle, given its radius

package main

import (
	"fmt"
)

func main() {
	const pi float32 = 3.414
	var radius float32 = 6.0
	fmt.Printf("Area: %v", pi*radius*radius)
}

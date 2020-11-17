package main

import "fmt"

func main() {
	var m int // declaration
	m = 20    // initialization

	var n int = 10 // declaration and initialization

	var o = 10 // inferred data-type

	p := 10

	fmt.Printf("%v %v %v %v: %T %T %T %T\n", m, n, o, p, m, n, o, p)
}

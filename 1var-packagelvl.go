package main

import "fmt"

var m int
var n int = 10

var o = 10

func main() {
	m = 10
	fmt.Printf("%v %v %v %T %T %T", m, n, o, m, n, o)
}

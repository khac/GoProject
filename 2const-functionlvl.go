package main

import "fmt"

const k int = 10
const j = 12
const o = 10

func main() {
	const f = 10
	const e string = "asd"
	fmt.Printf("%v %v %v %T %T %T\n", j, k, o, j, k, o)
	fmt.Printf("%v %v %T %T\n", e, f, e, f)
}

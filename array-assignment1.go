// declare and initialize an array of int

package main

import "fmt"

func main() {
	// Arrays have fixed length
	// once initialized their size cannot change

	var m [4]int // declaration
	m[0] = 20    // initialization
	m[1] = 21
	m[2] = 22
	m[3] = 23

	var n [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("%v %v: %T %T\n", m, n, m, n)
	fmt.Println(len(m), cap(m))
	fmt.Println(len(n), cap(n))
	fmt.Printf("%p %p", &m, &n)

}

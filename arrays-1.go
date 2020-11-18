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

	o := [3]int{11, 22, 33}
	p := [...]float32{1.23, 4.567, 8.1234}

	for i := 0; i < len(m); i++ {
		fmt.Println(i, m[i])
	}

	fmt.Println()

	for index, value := range n {
		fmt.Println(index, value)
	}

	fmt.Printf("%v %v %v: %T %T %T\n", m, n, o, m, n, o)
	fmt.Println(len(m), cap(m))
	fmt.Println(len(n), cap(n))
	fmt.Println(len(o), cap(o))
	fmt.Println(len(p), p)
}

package main

import "fmt"

func main() {
	// there is only for loop in go
	// do-while and while loop do not exist in golang
	for i := 0; i < 100; i = i + 7 {
		fmt.Println("", i)
	}
	fmt.Println()

}

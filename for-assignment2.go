// loop and print numbers from 0 to 100 in the increments of 2

package main

import "fmt"

func main() {
	// there is only for loop in go
	// do-while and while loop do not exist in golang
	for i := 0; i < 100; i = i + 2 {
		fmt.Println("", i)
	}
	fmt.Println()

}

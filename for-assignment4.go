// Print the sum of all the multiples of 3 and 5 below 1000

package main

import "fmt"

func main() {
	// there is only for loop in go
	// do-while and while loop do not exist in golang
	var sum = 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 && i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}

package main

import "fmt"

func main() {
	// there is only for loop in go
	// do-while and while loop do not exist in golang
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("", i)
	}
	fmt.Println("", sum)
	fmt.Println()

	j := 0
	sum = 0
	for ; j < 100; j = j + 21 {
		sum += j
		fmt.Println("", j)
	}
	fmt.Println("", sum)
	fmt.Println()

	j = 0
	sum = 0
	for ; ; j = j + 31 {
		if j > 100 {
			break
		}
		sum += j
		fmt.Println("", j)
	}
	fmt.Println("", sum)
	fmt.Println()

	j = 0
	sum = 0
	for {
		j += 41
		if j > 100 {
			break
		}
		sum += j
		fmt.Println("", j)
	}
	fmt.Println("", sum)
}

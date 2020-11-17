// given a number,
// if divisible by 3  then print foo
// if divisible by 5  then print bar
// if divisible by 15 then print foobar
// use switch instead of if

package main

import "fmt"

func main() {
	marks := 900

	switch {
	case marks%15 == 0:
		fmt.Println("foobar")
	case marks%3 == 0:
		fmt.Println("foo")
	case marks%5 == 0:
		fmt.Println("bar")
	}
}

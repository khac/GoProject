package main

import "fmt"

func main() {
	marks := 982

	switch {
	case marks < 35:
		fmt.Println("fail")
	case marks < 60:
		fmt.Println("pass")
	case marks < 75:
		fmt.Println("Pass with First class")
	case marks <= 100:
		fmt.Println("Pass with Distinction")
	default:
		fmt.Println("Enter valid marks")
	}

}

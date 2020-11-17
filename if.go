package main

import "fmt"

func main() {

	var marks int = 90
	if marks > 0 {
		fmt.Println("your marks are registered")
	}

	if marks < 35 {
		fmt.Println("fail")
	} else if marks < 60 {
		fmt.Println("Pass")
	} else if marks < 75 {
		fmt.Println("Pass with First Class")
	} else {
		fmt.Println("Pass with Distinction")
	}

	age := 20
	height := 152
	if age > 18 && height > 150 {
		fmt.Println("You are eligible")
	} else {
		fmt.Println("You are ineligible")
	}

	// && logical and
	// || logical or
	// ! logical not

	// if block
	// () are optional
	// {} are mandatory
	// first {, should be in the same line as the if/else

}

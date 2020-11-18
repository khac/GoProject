// Given slice of string, concatenate (join) them and form a single string, seperating them by a single space
// Input: “hello”,”world”
// Output: “hello world”

package main

import (
	"fmt"
	//"strings"
)

func main() {
	s := []string{"hello", "world"}

	fmt.Println(concat(s))
	fmt.Println(len(concat(s)))
}

func concat(s []string) string {
	// add your code here
	var cur string = ""
	for i:=0; i<len(s); i++ {
		cur += s[i]
		cur += " "
	}
	return cur[:len(cur)-1]
}

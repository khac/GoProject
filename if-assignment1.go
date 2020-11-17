// given a number, 
// if divisible by 3  then print foo 
// if divisible by 5  then print bar
// if divisible by 15 then print foobar

package main

import "fmt"

func main() {

	number := 300
	
	if(number%15==0){
		fmt.Println("foobar")
	} else if(number%3==0){
		fmt.Println("foo")
	} else if(number%5==0){
		fmt.Println("bar")
	}
}
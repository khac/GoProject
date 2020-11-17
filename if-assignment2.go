//task:
//if the value of the variable is between 1 and 9 print red
//10-19 print blue
//20-29 print yellow

package main

import "fmt"

func main() {
	number := 11

	if number >= 1 && number <= 9 {
		fmt.Println("red")
	} else if number >= 10 && number < 20 {
		fmt.Println("blue")
	} else if number >= 20 && number < 30 {
		fmt.Println("yellow")
	}

}

// create a function with person struct input (with firstname and lastname fields). 
// the function should derive the full name and return it,

package main

import (
	"fmt"
)

type person struct {
    firstname string
    lastname  string
}

func main() {
	p := person{}
	p.firstname = "Abcd"
	p.lastname = "xyz"
	
	fmt.Println(fullName(p))
}

func fullName(s person) string {
	
	return s.firstname+" "+s.lastname
}

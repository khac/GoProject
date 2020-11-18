// create a "person" struct with fields "firstname", "lastname". 
// add a method called "fullname" to return the full name based on its "firstname" and "lastname" fields

package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName    string	
}

func (p *person) fullName() string {
    return p.firstName+" "+p.lastName
}

func main() {
	p1 := person{"ASDA", "MNBMCNB"}
	fmt.Println(p1)
	fmt.Println(p1.fullName())
}

// embed "person" struct in "personal" struct
// print "firstname" field value from "personal" struct's variable

package main

import (
	"fmt"
)

type personal struct {
	address string
	city    string
	state   string
	person  person
}

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{"ASDA POIPOI", 121}
	w1 := personal{"123 Blvd J Street", "asda", "asdad", p1}

	fmt.Println(w1.person.name)
	fmt.Println(w1.person.age)
}

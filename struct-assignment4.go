// create an "homeaddress" struct with fields address, city, state, country.
// create an "workaddress" struct with fields address, city, state, country.
// embed these two structs in the "person" struct for the home address and work address sections

package main

import (
	"fmt"
)

type homeaddress struct {
	address string
	city    string
	state   string
	country string
}

type workaddress struct {
	address string
	city    string
	state   string
	country string
}

type person struct {
	name    string
	age     int
	wrkaddr workaddress
	hmeaddr homeaddress
}

func main() {
	h1 := homeaddress{"123 Ave Apt. 21", "xaas", "Oak", "WWKJA"}
	w1 := workaddress{"123 Blvd J Street", "asda", "asdad", "asda"}
	p1 := person{"ASDA POIPOI", 121, w1, h1}
	fmt.Println(p1)

}

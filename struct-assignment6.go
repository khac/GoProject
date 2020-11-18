// create an "address" struct with fields address, city, state, country. 
// add the "address" struct in the "person" struct for the home address and work address sections

package main

import (
	"fmt"
)

type address struct {
	address string
	city    string
	state   string
	country  string
}

type person struct {
	name string
	age  int
	homeaddr address
	workaddr address
}

func main() {
	workAdd := address{"123 Blvd J Street", "Csda", "AS", "AAS"}
	homeAdd := address{"111 St. #123", "Aksalmn", "XS", "AAS"}
	p1 := person{"ASDA MNBMCNB", 21, homeAdd, workAdd}
	fmt.Println(p1)
}

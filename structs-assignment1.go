// to create a "person" type struct with fields "firstname" and "lastname"

package main

import (
	"fmt"
)

type person struct{
	firstname string
	lastname string
}

func main() {
	m := person{firstname:"xyz", lastname:"abc"}
	fmt.Println(m)
}
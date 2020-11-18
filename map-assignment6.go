//Create a map
//add one key value pair: 'Contact' as key and '000-000-0000' as value
//now, update a value of 'Contact' key to '123-456-7890'

package main

import (
	"fmt"
)

func main() {
	m := map[string]string{}
	m["Contact"] = "000-000-0000"
	m2 := update(m, "Contact", "123-456-7890")
	fmt.Println(m2)
}

func update(m map[string]string, k string, v string) map[string]string {
	// add your code here
	m[k] = v
	return m
}

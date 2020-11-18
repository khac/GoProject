//Create a map
//add one key value pair: 'Contact' as key and '000-000-0000' as value
//now, delete the pair with key as Contact

package main

import (
	"fmt"
)

func main() {
	m := map[string]string{}
	m["Contact"] = "123-456-7890"

	m = delKey(m, "Contact")
	fmt.Println(m)
}

func delKey(m map[string]string, k string) map[string]string {
	// add your code here
	delete(m, k)
	return m
}

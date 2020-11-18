//Intialise map with key as string and value as int


//Create a map using make


//Print a map with key as string and value as string without assigning any values

package main

import (
	"fmt"
)

func main() {
	exampleMap := make(map[string]int)
	exampleMap["one"] = 1
	exampleMap["two"] = 2
	exampleMap["three"] = 3
	exampleMap["four"] = 4
	exampleMap["five"] = 5
	exampleMap["six"] = 6
	for k,v := range(exampleMap){
		fmt.Println(k,v)
	}
}

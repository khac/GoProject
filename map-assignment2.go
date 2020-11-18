//Create a map using make


//Print a map with key as string and value as string without assigning any values

package main

import (
	"fmt"
)

func main() {
	exampleMap := make(map[string]string)
	for k,v := range(exampleMap){
		fmt.Println(k,v)
	}
}

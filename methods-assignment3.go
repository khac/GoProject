// create a "square" struct with "side" as field. 
// and add "area" method that derives the area from its "side" field

package main

import (
	"fmt"
)

type square struct {
	side float32
}

func (p *square) area() float32 {
    return p.side*p.side
}

func main() {
	p1 := square{11.1}
	fmt.Println(p1)
	fmt.Println(p1.area())
}

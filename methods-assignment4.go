// create a "rectangle" struct with "length", "width" as fields. and add "area" method that derives the area from its "length", "width" field

package main

import (
	"fmt"
)

type rectangle struct {
	length float32
	width  float32
}

func (p *rectangle) area() float32 {
	return p.length * p.width
}

func main() {
	p1 := rectangle{11.1, 22.2}
	fmt.Println(p1)
	fmt.Println(p1.area())
}

// create a "circle" struct with "radius" as field. 
// and add "area" method that derives the area from its "radius" field


package main

import (
	"fmt"
)

type circle struct {
	radius float32
}

func (p *circle) area() float32 {
    return 3.14*p.radius*p.radius
}

func main() {
	p1 := circle{11}
	fmt.Println(p1)
	fmt.Println(p1.area())
}

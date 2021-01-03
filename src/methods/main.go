// Go by example - methods
package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	result := rect{
		width:  10,
		height: 5,
	}

	fmt.Println("area:", result.area())
	fmt.Println("perim:", result.perim())

	rp := &result
	fmt.Println(rp) // print out rp to learn what assigned to rp

	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

}

// Interface - example 1
package main

import (
	"fmt"
	"strings"
)

type shape interface {
	area() int
	perim() int
}

type rectangle struct {
	w, l int // width & length
}

type square struct {
	s int // side
}

func (a *rectangle) area() int {
	return a.w * a.l
}

func (a *rectangle) perim() int {
	return 2 * (a.w + a.l)
}

func (a *square) area() int {
	return a.s * a.s
}

func (a *square) perim() int {
	return 4 * a.s
}

func info(s shape) {
	fmt.Printf("Area() = %-5d perim() = %d\n", s.area(), s.perim())
}

func totalArea(shapes ...shape) int {
	var totalArea int

	for _, s := range shapes {
		totalArea += s.area()
	}
	return totalArea
}

func main() {

	r1 := rectangle{20, 30}
	fmt.Printf("r1.area()\t= %-5d\n", r1.area())
	fmt.Printf("r1.perim()\t= %-5d\n", r1.perim())

	r2 := rectangle{15, 25}
	fmt.Printf("r2.area()\t= %d\n", r2.area())
	fmt.Printf("r2.perim()\t= %d\n", r2.perim())

	s1 := square{35}
	fmt.Printf("s1.area()\t= %d\n", s1.area())
	fmt.Printf("s1.perim()\t= %d\n", s1.perim())

	s2 := square{25}
	fmt.Printf("s2.area()\t= %d\n", s2.area())
	fmt.Printf("s2.perim()\t= %d\n", s2.perim())

	fmt.Println("\n", strings.Repeat("=", 35))
	info(&r1)
	info(&r2)
	info(&s1)
	info(&s2)

	fmt.Println("\n", strings.Repeat("=", 35))
	fmt.Printf("Total area = %d\n", totalArea(&r1, &r2, &s1, &s2))

	shapes := []shape{&r1, &r2, &s1, &s2}

	totalArea := 0
	for _, s := range shapes {
		totalArea += s.area()
	}

	fmt.Printf("Total Area = %d\n", totalArea)
}

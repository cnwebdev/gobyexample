// variable, constant and literals
package main

import "fmt"

func main() {

	var k1 uint8 = 20

	const pi = 3.14159

	k1 /= 2

	fmt.Println(25, pi, k1)

	// variable are statically typed
	// K1 = "changing type is not allowed." // compiler error

	// Constant variablies
	const (
		a = 2      // a = 2
		b          // b = 2
		c          // c = 2
		d = a * 10 // 20
		e          // 20
		f          // 20
	)
	fmt.Println(a, b, c, d, e, f)

	// const with iota
	const (
		a2 = iota + 1 // a2 = 1
		b2            // b2 = 2
		c2            // c2 = 3
		d2 = a2 * 10  // d2 = 10
		e2            // e2 = 10
		f2 = iota     // f2 = 5
	)
	fmt.Println(a2, b2, c2, d2, e2, f2)

	// const with iota and bit shifting
	const (
		a3 = iota + 1 // a3 = 1 = 00000001
		b3 = a3 << 1  // b3 = 2 = 00000010
		c3 = b3 << 1  // c3 = 4 = 00000100
		d3 = c3 << 1  // d3 = 8 = 00001000
		e3 = d3 << 1  // e3 = 16 = 00010000
	)
	fmt.Println(a3, b3, c3, d3, e3)
}

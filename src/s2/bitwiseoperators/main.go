// Bitwise operators
package main

import "fmt"

func main() {

	var x uint8 = 2

	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)

	x = x << 1
	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)

	x = x >> 2
	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)

	x = 4 | 2 // 0000 0100 | 0000 0010 = 0000 0110
	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)

	x = 4 & 2 // 0000 0100 & 0000 0010 = 0000 0000
	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)

	y := ^4
	fmt.Printf("%8d %#8o %#8x \t %08b\n", y, y, y, y)

}

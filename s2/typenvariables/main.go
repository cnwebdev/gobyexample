// type and variables example
package main

import "fmt"

func main() {
	// declare variable i without assign/initialize a value
	var i int
	fmt.Println("i =", i) // i is 0

	// delcare and assign 2
	j := 2
	fmt.Printf("j = %d, j type = %T\n", j, j) // %d = print integer; %T = print type

	// k1 is uint8 = only positive value 0 - 255
	var k1 uint8 = 20
	fmt.Printf("k1 = %d, k1 type = %T\n", k1, k1)

	// uint32 = 0 to 4294967295
	var a, b, c uint32
	a = 321325
	b = 424521
	c = a * b
	fmt.Println(c) // 3265224149 < uint32 = 0 to 4294967295
}

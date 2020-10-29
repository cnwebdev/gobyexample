// struct example 2 (Single type)
package main

import "fmt"

func main() {
	type custom float64
	var total custom

	total = 55
	fmt.Printf("%.2f %T \n", total, total)

	var total2 float64

	// total2 = total // compiler error
	total2 = float64(total)
	fmt.Printf("%.2f %T \n", total2, total2)

	// total == total2 // type miss match error

}

// Go by example - functions
package main

import "fmt"

func main() {
	res := plus(3, 8)
	fmt.Printf("Sum of 3 and 8 = %d\n", res)

	res = plusPlus(5, 3, 9)
	fmt.Printf("Sum of 5, 3, and 9 = %d\n", res)
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// Go by example - multiple return values
package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {

	// get return values and save to variable a and b
	a, b := vals()
	fmt.Println(a, b)

	// ignore one return value
	_, c := vals()
	fmt.Println(c)
}

// Unicode part 1
package main

import "fmt"

func main() {
	ascii := 'a'
	world := '世'
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", world)
	fmt.Printf("%d %[1]c %[1]q\n", newline)
}

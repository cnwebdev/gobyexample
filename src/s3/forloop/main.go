// For loop statements - example 1
package main

import "fmt"

func main() {

	// repeat the same operations
	fmt.Println(2, 20)
	fmt.Println(3, 30)
	fmt.Println(4, 40)

	fmt.Println()
	i := 2

	// repeat the loop as loong as i < 5, exit loop when i = 5
	for i < 5 {
		fmt.Println(i, i*10)
		i++
	}
}

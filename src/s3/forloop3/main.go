// For loop statements - example 3
package main

import "fmt"

func main() {
	i := 1

	// loop until i == 8
	for {

		if i == 8 {
			break
		}

		if i == 4 || i == 7 {
			i++
			continue
		}

		fmt.Print(i, " ")
		i++
	}
}

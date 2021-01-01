// for loop statements - example 6
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("\nMultiplication table = %d ** ", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%-3d ", i*j)
		}
	}
}

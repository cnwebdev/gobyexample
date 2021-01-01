// map - example 1
package main

import "fmt"

func main() {
	values := []string{"ABC", "ACB", "BAC", "BCA", "CBA"}

	factor := []int{100, 10, 1}

	key := 0

	for i := range values {

		bytes := []byte(values[i])
		f := 0
		key = 0
		for j := range bytes {
			fmt.Print(bytes[j], " ")
			key += int(bytes[j]) * factor[f]
			f++
		}
		fmt.Printf(" (key: %d) \n", key)
	}
}

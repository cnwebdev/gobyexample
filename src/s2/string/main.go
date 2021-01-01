// Working with strings
package main

import "fmt"

func main() {
	fmt.Println(len("Good morning America"))
	fmt.Println("Good morning"[5])
	fmt.Println(string("Good morning"[5]))
	fmt.Println()

	s := "The United States of America"

	fmt.Println(s[:10])
	fmt.Println(s[10:])
	fmt.Println(s[4:10])
	fmt.Println(s[:])

	fmt.Println(s[:10] + s[10:])

	// s[5] = "B" // compile error

	s += "'s Flag"
	fmt.Println(s)
}

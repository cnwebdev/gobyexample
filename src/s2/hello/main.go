// Hello world program

package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("Hello, 世界!") // Hello 世界 (世界 means world in japanese)

	var firstName = "Chris" // type string
	lastName := `Nguyen`    // type string

	const (
		birthYear = 1975 // type int
		currYear  = 2020 // type int
	)

	fmt.Printf("%s %s\n", firstName, lastName)

	fmt.Printf("Age: %d\n", currYear-birthYear)
	fmt.Println("A\tB\a") // "\a" will play a sound

	n := 5
	fmt.Printf("%[1]d + %[1]d = %[1]d\n", n, n+n)
}

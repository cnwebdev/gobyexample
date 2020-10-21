// Hello world program

package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	var firstName = "Chris" // type string
	lastName := `Nguyen`    // type string

	const (
		birthYear = 1975 // type int
		curYear   = 2020 // type int
	)

	fmt.Printf("%s %s\n", firstName, lastName)

	fmt.Printf("Age: %d\n", curYear-birthYear)
	fmt.Println("A\tB\a")

	n := 5
	fmt.Printf("%[1]d + %[1]d = %[1]d\n", n, n+n)
}

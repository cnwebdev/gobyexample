// Character escapes
package main

import "fmt"

func main() {

	fmt.Println("col1\tcol2\tcol3")
	fmt.Print("--------------------------------\n")
	fmt.Println("item11\titem12\titem13")
	fmt.Println("item21\titem22")

	fmt.Printf("\n%s\t%s\t%s", "col1", "col2", "col3")
	fmt.Print("\n-------------------------------")
	fmt.Printf("\n%s\t%s\t%s", "item11", "item12", "item13")
	fmt.Printf("\n%s\t%s\t%s\n", "item21", "", "item23")
}

// switch stattements - example 7
package main

import "fmt"

func main() {
	var x interface{}

	x = mult(10)

	switch i := x.(type) {
	case int:
		fmt.Printf("%T %v\n", i, i)
	case float64:
		fmt.Printf("%T %v\n", i, i)
	case bool, string:
		fmt.Println("type is bool of string")
	case func(int) float64:
		fmt.Printf("%T\n", i)
	case nil:
		fmt.Println("x is nil")
	default:
		fmt.Println("don't know the type")
	}
}

func mult(i int) float64 {
	return float64(i) * 1.5
}

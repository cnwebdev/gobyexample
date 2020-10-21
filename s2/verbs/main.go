package main

import "fmt"

func main() {
	var i uint8 = 128 // largist value for uint8
	var f float32 = 233.5
	var b bool = true
	var msg string = "United States of America"

	type myString string

	var newMsg myString = "The United States Flag has 50 stars and 13 strips"

	fmt.Printf("%d %v %T \n", i, i, i)
	fmt.Printf("%5.1f %.2f %g %T\n", f, f, f, f)
	fmt.Printf("%t %v %T\n", b, b, b)
	fmt.Printf("%s %T \n", msg, msg)
	fmt.Printf("%v %T \n", newMsg, newMsg)

}

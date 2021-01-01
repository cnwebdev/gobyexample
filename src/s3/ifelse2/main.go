// if statements example 2
package main

import "fmt"

func main() {
	seasonNo := 3

	if seasonNo == 1 {
		fmt.Println("Spring -", seasonNo)
	} else if seasonNo == 2 {
		fmt.Println("summer -", seasonNo)
	} else if seasonNo == 3 {
		fmt.Println("fall -", seasonNo)
	} else if seasonNo == 4 {
		fmt.Println("Winter -", seasonNo)
	} else {
		fmt.Println("a brand new season - ", seasonNo)
	}
}

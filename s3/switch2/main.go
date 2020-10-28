// switch statements - example 2
package main

import "fmt"

func main() {
	for {

		var seasonNo int
		fmt.Print("Enter a number:")
		fmt.Scanf("%d", &seasonNo)

		switch seasonNo {
		case 1:
			fmt.Println("Spring - ", seasonNo)
		case 2:
			fmt.Println("summer - ", seasonNo)
		case 3:
			fmt.Println("fall - ", seasonNo)
		case 4:
			fmt.Println("winter - ", seasonNo)
		case 5:
			return
		default:
			fmt.Println("a new season - ", seasonNo)
		}
	}
}

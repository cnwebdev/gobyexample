// switch statements - example 3
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
		case 3:
		case 4:
			fmt.Println("summer - ", seasonNo)
		case 5:
			return
		default:
			fmt.Println("a new season - ", seasonNo)
		}
	}
}

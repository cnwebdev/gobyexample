// switch statements - example 4
package main

import "fmt"

func main() {
	for {
		var seasonNo int
		fmt.Print("Enter a number:")
		fmt.Scanf("%d", &seasonNo)

		switch seasonNo {
		case 1:
			fmt.Println("spring - ", seasonNo)
		case 2, 3, 4:
			fmt.Println("summer - ", seasonNo)
			fmt.Println("fall - ", seasonNo)
			fmt.Println("winter - ", seasonNo)
		case 5:
			fmt.Println("Exit!")
			return
		default:
			fmt.Println("a new season - ", seasonNo)
		}
	}
}

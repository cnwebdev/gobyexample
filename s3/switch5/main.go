// switch statements - example 5
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var season string

	for season != "end" {
		fmt.Print("Enter your faverite season (type end or exit): ")
		fmt.Scanf("%s ", &season)

		s := strings.ToLower(season)
		switch s {
		// s := strings.ToLower(season)
		case "spring":
			fmt.Fprintf(os.Stdout, "You entered %s\n", s)
		case "summer":
			fmt.Println("You Entered", s)
			break
			fmt.Println("This line won't be reached", s)
		case "fall":
			fmt.Println("You Entered", s)
		case "winter":
			fmt.Println("You Entered", s)
		case "exit":
			fmt.Println("Exit!")
			return
		default:
			fmt.Println("a new season - ", s)
		}
	}
}

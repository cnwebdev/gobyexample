// for loop statements - example 5
package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Press CTRL+C to stop the program")

	i := 2
	for {
		fmt.Println(i, " ")
		i += 2
		time.Sleep(500 * time.Millisecond) // wait for .5 second
	}
}

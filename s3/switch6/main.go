// switch statements - example 6
package main

import "fmt"

func main() {
	sportID := 2
	isYoung := true

	switch {
	case sportID == 1 && isYoung:
		fmt.Println("playing soccer ...")
	case sportID == 1 && !isYoung:
		fmt.Println("watching soccer ...")
	case sportID == 2 && isYoung:
		fmt.Println("playing chess ...")
	case sportID == 2 && !isYoung:
		fmt.Println("still can play chess ...")
	default:
		fmt.Println("We'll discuss later ...")
	}
}

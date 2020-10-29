// struct pointers
package main

import "fmt"

func main() {
	player1 := struct {
		name, sport string
		age         int
		rank        float64
	}{"Duan House", "Tennis", 39, 4.5}

	fmt.Println("Duan's", player1)

	player2 := struct {
		name, sport string
		age         int
	}{
		age:   30,
		sport: "Soccer",
		name:  "J.C.",
	}
	fmt.Println("J.C", player2)
}

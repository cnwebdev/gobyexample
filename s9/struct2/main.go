// structexample 1
package main

import "fmt"

type player struct {
	name  string
	sport string
	age   int
	rank  float64
}

func main() {

	// declare a new player base on type player struct
	kevin := player{name: "Kevin", sport: "Tennis", age: 39, rank: 4.5}
	// Print kevin player data
	fmt.Println("Kevin's profile:", kevin)
	// Access specific fields
	fmt.Printf("name: %s age: %d\n", kevin.name, kevin.age)

	duan := player{
		age:   36,
		sport: "Tennis",
		name:  "Duan",
		rank:  4.5,
	}
	// Print kevin player data
	fmt.Println("Duan's profile:", duan)
	// Access specific fields
	fmt.Printf("name: %s age: %d\n", duan.name, duan.age)

	test := kevin == duan
	fmt.Println("Kevin and Duan are the same?", test)

	// players with zero value
	var chris player
	fmt.Printf("chris: %+v\n", chris)

	// Assign value to fields
	chris.name = "Chris"
	chris.age = 50
	chris.sport = "Tennis"
	chris.rank = 4.0

	// new player with not all fields initialized
	hank := player{
		name:  "Hank Vu",
		sport: "Tennis",
	}
	fmt.Printf("Hank Vu %+v\n", hank)

	// create new player with new
	john := new(player)
	fmt.Printf("john: %+v\n", john)
	john.name = "Johnathan"
	john.sport = "Tennis"
	john.age = 39
	fmt.Printf("john: %+v\n", john)
}

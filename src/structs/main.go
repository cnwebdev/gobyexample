// Go by example - structs
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	p1 := person{
		name: "Bob",
		age:  20,
	}
	fmt.Println(p1)

	p2 := person{
		name: "Alice",
		age:  30,
	}
	fmt.Println(p2)

	p3 := person{
		name: "Fred",
	}
	fmt.Println(p3)

	/*
		p4 := person{
			name: "Ann",
			age:  40,
		}
		fmt.Println("p4", p4)
	*/

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp)

	sp.age = 51
	sp.name = "Chris"
	fmt.Println(s.age, s.name)
	fmt.Println(sp)
	fmt.Println(s)
}

func newPerson(name string) *person {

	p := person{
		name: name,
	}

	p.age = 42
	return &p
}

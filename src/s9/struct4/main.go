// struct - example 4
package main

import (
	"fmt"
	"time"

	"github.com/fxtlabs/date"
)

type employee struct {
	name, job, lastLoggedIn string
	DOB                     date.Date
}

func main() {

	var emp employee

	emp.name = "Nick"
	emp.job = "Go Programmer"
	emp.lastLoggedIn = time.Now().Format(time.RFC850)

	emp.DOB = date.Today()

	fmt.Println(emp)

	emp.name = "Mike"
	fmt.Println("Mike", emp)

	// This will make a copy of emp whatever change in test will not effect emp.
	test := emp

	test.name = "Kim"
	fmt.Println("Kim", emp, test)

	// Test really chage emp, we can assign pointer variable.
	p := &emp

	p.name = "Chris"
	emp.job = "Go Expert"
	fmt.Println(*p)

	fmt.Printf("%x %x %x\n", &emp.name, &p.name, &test.name)

}

// Go by exmaple - slices
package main

import "fmt"

func main() {
	// declare s slice with lenght = 3 and cap = 3
	s := make([]string, 3)
	fmt.Printf("emp: len:%d, cap:%d\n", len(s), cap(s))

	// assign string "a", "b", "c" to c slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set", s)
	fmt.Println("get", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// assign c new slice with same length of s slice
	c := make([]string, len(s))

	// compy the content of slice s to slice c
	copy(c, s)
	fmt.Printf("s lice: %q\n", s)
	fmt.Printf("c lice: %q\n", c)

	// slicing s slice for 2 to 4 elements
	l := s[2:5]
	fmt.Println("sl1:", l)

	// slicing s slice for 0 to 4
	l = s[0:5]
	fmt.Println("sl2:", l)

	// slicing s slice for 2 - end of the the slice
	l = s[2:]
	fmt.Println("sl3:", l)

	// declare and assign slice t
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// declare multi-dimentional slice
	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

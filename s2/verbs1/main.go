// verbs1 exmaple
package main

import "fmt"

func main() {

	var k1 uint8 = 20

	const pi = 3.14

	fmt.Printf("%f, %T\n", pi, pi)           // 3.140000, float64
	fmt.Printf("%5.3f %.2f \n", pi, 213.438) // 3.140 213.44

	fmt.Println(k1+15, k1-12, k1*2, k1/2, 5%2) // 35, 8, 40, 10, 1
	k1++
	fmt.Println(k1) // 21

	k1 += 10
	fmt.Println(k1) // 31

	fmt.Println(5/2, 5.0/2.0) // 2, 2.5
}

// Go by example - constant
package main

import (
	"fmt"
	"math"
)

const s = "Constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	fmt.Println(n)

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

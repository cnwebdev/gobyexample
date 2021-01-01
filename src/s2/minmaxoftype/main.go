// Minium and maximum values of basic types
package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("%-15s %-30d %-30d\n", "int8", math.MinInt8, math.MaxInt8)
	fmt.Printf("%-15s %-30d %-30d\n", "int16", math.MinInt16, math.MaxInt16)
	fmt.Printf("%-15s %-30d %-30d\n", "int32", math.MinInt32, math.MaxInt32)
	fmt.Printf("%-15s %-30d %-30d\n", "int64", math.MinInt64, math.MaxInt64)
	fmt.Println()

	fmt.Printf("%-15s %-30d %-30d\n", "uint8", 0, math.MaxUint8)
	fmt.Printf("%-15s %-30d %-30d\n", "uint16", 0, math.MaxUint16)
	fmt.Printf("%-15s %-30d %-30d\n", "uint32", 0, math.MaxUint32)
	fmt.Printf("%-15s %-30d %-30d\n", "uint64", 0, uint64(math.MaxUint64))

	fmt.Println()
	fmt.Printf("%-60v\n", math.MaxFloat32)
	fmt.Printf("%-60v\n", math.MaxFloat64)
}

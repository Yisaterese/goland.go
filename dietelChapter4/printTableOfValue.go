package main

import (
	"fmt"
	"math"
)

func printTableOfValues() {
	var (
		n2 = 1
		n3 = 1
		n4 = 1
	)

	fmt.Println("N\tN2\tN3\tN4")
	for count := 1; count <= 5; count++ {
		n2 = int(math.Pow(float64(count), float64(2)))
		n3 = int(math.Pow(float64(count), float64(3)))
		n4 = int(math.Pow(float64(count), float64(4)))
		fmt.Printf("%d\t%d\t%d\t%d\n", count, n2, n3, n4)
	}
}

func main() {
	printTableOfValues()
}

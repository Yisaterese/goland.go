package main

import "fmt"

func main() {
	rows := 8
	cols := 8

	for i := 0; i < rows; i++ {
		// If the row is odd, start the inner loop from 1
		// Otherwise, start from 0
		start := 0
		if i%2 != 0 {
			start = 1
		}

		for j := start; j < cols; j += 2 {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

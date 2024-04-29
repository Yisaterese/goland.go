package main

import (
	"fmt"
)

func checkBoardPattern() {
	for z := 1; z < 8; z++ {
		for i := 1; i < 8; i++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
func main() {
	checkBoardPattern()
}

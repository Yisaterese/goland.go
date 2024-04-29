package main

import "fmt"

func checkBoardPattern() {
	for i := 1; i < 8; i++ {
		for j := 1; j <= 8; j++ {
			if i%2 == 0 {
				fmt.Print("* ")
			} else {
				fmt.Print(" *")
			}
		}
		fmt.Println()
	}
}
func main() {
	checkBoardPattern()

}

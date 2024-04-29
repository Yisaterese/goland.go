package main

import "fmt"

func displayARightAngleTriangle(userInput int) {
	for row := 1; row <= userInput; row++ {
		for col := 1; col <= row; col++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
func main() {
	displayARightAngleTriangle(5)
}

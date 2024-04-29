package main

import "fmt"

func findTheLargestNumber() {
	var (
		largest   = 0
		userInput = 0
	)
	fmt.Println("Enter your number:")
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		return
	}
	for counter := 1; counter < 10; counter++ {
		fmt.Println("Enter your number:")
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			return
		}
		if userInput > largest {
			largest = userInput
		}
	}
	//return largest
	fmt.Println(largest)
}

func main() {
	findTheLargestNumber()
}

package main

import "fmt"

func addWithOutUsingAddition(number1, number2 int) int {
	return number1 - -number2
}
func main() {
	fmt.Println(addWithOutUsingAddition(1, 1))
}

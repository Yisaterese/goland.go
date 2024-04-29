package main

import "fmt"

// exercise 4.6
func sumOfIntegers() {
	var (
		total = 0
		x     = 1
	)
	for ; x < 11; x++ {
		total += x
	}
	fmt.Println(total)
}
func main() {
	sumOfIntegers()

}

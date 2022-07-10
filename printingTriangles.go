package main

import "fmt"

func main() {
	printAscendingTriangle(10)
	printDescendingTriangle(10)

}
func printAscendingTriangle(numberOfRow int) {
	for i := 0; i < numberOfRow; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()

	}

}
func printDescendingTriangle(numberOfRow int) {
	for i := numberOfRow; i > 0; i-- {
		for j := 0; j < i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

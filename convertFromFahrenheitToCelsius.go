package main

import "fmt"

func main() {
	fmt.Print("Enter a number in fahrenheit: ")
	var userInput float64
	fmt.Scanf("%f", &userInput)
	fahrenheitToCelsius(userInput)

}
func fahrenheitToCelsius(F float64) {
	C := (F - 32) * 5 / 9

	fmt.Println(C)

}

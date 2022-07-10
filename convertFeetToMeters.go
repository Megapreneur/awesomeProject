package main

import "fmt"

func main() {
	fmt.Print("Enter a number in feet: ")
	var userInput float64
	fmt.Scanf("%f", &userInput)
	feetToMeters(userInput)
}
func feetToMeters(feet float64) {
	m := feet * 0.3048
	fmt.Println(feet, "feet is equal to ", m, "meters")

}

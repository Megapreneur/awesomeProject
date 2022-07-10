package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var userInput int
	fmt.Scanf("%d", &userInput)
	printingEvenNumbers(userInput)
}
func printingEvenNumbers(number int) {
	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is an even number")
		} else {
			fmt.Println(i, " is an odd number")
		}

	}

}

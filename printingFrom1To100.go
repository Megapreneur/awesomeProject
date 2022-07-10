package main

import "fmt"

func main() {
	fmt.Print("Enter the limiting number: ")
	var userInput int
	fmt.Scanf("%d", &userInput)
	printingNumbers(userInput)

}
func printingNumbers(limit int) {
	i := 1
	for i <= limit {
		fmt.Println(i)
		i++
	}
}
